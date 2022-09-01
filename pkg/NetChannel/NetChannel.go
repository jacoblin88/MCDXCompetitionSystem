package NetChannel

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PangolinoLab/MCDXCompetitionSystem/pkg/KeyMutex"
)

var constValue = struct {
	logOutFile *os.File
}{
	logOutFile: os.Stdout,
}

type NetChannel struct {
	sync.WaitGroup
	ip              string
	port            int
	listenHandler   net.PacketConn
	dialConn        net.Conn
	myPrivateKey    *rsa.PrivateKey
	targetPublicKey *rsa.PublicKey
	timeout         float32
	tryTimes        int
	blockSzLimit    int
}

type RecvCallBackFunc func(msg []byte, msgSz int) (respMsg []byte, err error)

var resourceLock = KeyMutex.NewHashed(83)

func New(ip string, port int, timeout float32, tryTimes int, myKeyJson string, targetKeyJson string) (*NetChannel, error) {
	logger := log.New(constValue.logOutFile, "n.New "+" ", log.Ldate|log.Ltime|log.Lshortfile)
	var n NetChannel
	tmp, err := n.rsaKeySetup([]byte(myKeyJson), true)
	if err != nil {
		logger.Panic(err)
		return nil, err
	}
	n.myPrivateKey = tmp
	tmp, err = n.rsaKeySetup([]byte(targetKeyJson), false)
	if err != nil {
		logger.Panic(err)
		return nil, err
	}
	n.targetPublicKey = &tmp.PublicKey
	n.timeout = timeout
	n.tryTimes = tryTimes
	n.blockSzLimit = 1024
	n.ip = ip
	n.port = port

	if n.timeout < 0 {
		resourceLock.LockKey(n.ip + ":" + strconv.Itoa(n.port))
		n.listenHandler, err = net.ListenPacket("udp", ip+":"+strconv.Itoa(port))
		if err != nil {
			logger.Panic(err)
		}
		n.dialConn = nil
	} else {
		n.listenHandler = nil
		n.dialConn, err = net.Dial("udp", ip+":"+strconv.Itoa(port))
		if err != nil {
			logger.Panic(err)
		}
	}

	return &n, nil
}

func (n *NetChannel) Send(msgStr []byte, msgSz int) (int, error) {
	logger := log.New(constValue.logOutFile, "n.Send "+n.ip+":"+strconv.Itoa(n.port)+" ", log.Ldate|log.Ltime|log.Lshortfile)
	if n.timeout < 0 {
		return 0, errors.New("send can not use when in server mode")
	} else {
		encMsg, err := n.rsaEncrypt(n.targetPublicKey, msgStr[:msgSz])
		if err != nil {
			logger.Println(err)
			return 0, err
		}
		if len(encMsg) > n.blockSzLimit {
			logger.Println(errors.New("msg too long"))
			return 0, errors.New("msg too long")
		}
		sentSz, err := n.dialConn.Write(encMsg)
		if sentSz != len(encMsg) && err == nil {
			err = errors.New("only sent part of msg")
			return 0, err
		} else if err != nil {
			return 0, err
		}
		return msgSz, err
	}
}

func (n *NetChannel) Recv(f RecvCallBackFunc) ([]byte, error) {
	logger := log.New(constValue.logOutFile, "n.Recv "+n.ip+":"+strconv.Itoa(n.port)+" ", log.Ldate|log.Ltime|log.Lshortfile)
	buff := make([]byte, n.blockSzLimit)
	var msg []byte
	var buffLen int
	if n.timeout < 0 {
		// if err := n.listenHandler.SetReadDeadline(time.Now().Add(time.Second * time.Duration(30))); err != nil {
		// 	logger.Panic(err)
		// }
		buffLen, addr, err := n.listenHandler.ReadFrom(buff)
		if err != nil && (!strings.Contains(err.Error(), "use of closed network connection")) {
			logger.Println(err)
			return buff[:buffLen], err
		}
		msg, err = n.rsaDecrypt(*n.myPrivateKey, buff[:buffLen])
		if err != nil {
			logger.Println(err)
			return msg, err
		}
		n.Add(1)
		go n.triggerCallBackFunc(msg, f, addr)
	} else {
		var err error
		for i := 0; i < n.tryTimes; i++ {
			if err := n.dialConn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(n.timeout))); err != nil {
				logger.Println(err)
			}
			buffLen, err = n.dialConn.Read(buff)
			if err != nil {
				logger.Println(err)
				continue
			}
			msg, err = n.rsaDecrypt(*n.myPrivateKey, buff[:buffLen])
			if err != nil {
				logger.Println(err)
				continue
			}
			n.Add(1)
			go n.triggerCallBackFunc(msg, f, nil)
			break
		}
		return msg, err
	}
	return msg, nil
}

func (n *NetChannel) triggerCallBackFunc(buff []byte, f RecvCallBackFunc, addr net.Addr) {
	logger := log.New(constValue.logOutFile, "n.triggerCallBackFunc "+n.ip+":"+strconv.Itoa(n.port)+" ", log.Ldate|log.Ltime|log.Lshortfile)
	defer n.Done()
	if addr != nil {
		respBuff, err := f(buff, len(buff))
		if err != nil {
			logger.Println("err", err)
		}
		if respBuff == nil {
			return
		}
		encMsg, err := n.rsaEncrypt(n.targetPublicKey, respBuff)
		if err != nil {
			logger.Println(err)
		}
		if len(encMsg) > n.blockSzLimit {
			logger.Println(err)
		}
		n.listenHandler.WriteTo(encMsg, addr)
	} else {
		respBuff, err := f(buff, len(buff))
		if err != nil {
			logger.Println("respBuff", respBuff, "err", err)
		}
		if respBuff == nil {
			return
		}
		encMsg, err := n.rsaEncrypt(n.targetPublicKey, respBuff)
		if err != nil {
			logger.Println(err)
		}
		if len(encMsg) > n.blockSzLimit {
			logger.Println(err)
		}
		sentSz, err := n.dialConn.Write(encMsg)
		if sentSz != len(encMsg) && err == nil {
			err = errors.New("only sent part of msg")
			logger.Println(err)
		} else if err != nil {
			logger.Println(err)
		}
		return
	}
}

func (n *NetChannel) SetDestination(interface{}) error {
	return errors.New("unimplemented features")
}

func (n *NetChannel) GetDestination() (interface{}, error) {
	return nil, errors.New("unimplemented features")
}

func (n *NetChannel) CheckStatus() error {
	if n.timeout < 0 && n.listenHandler == nil {
		return errors.New("listen handler non-exist")
	} else if n.timeout >= 0 && n.dialConn == nil {
		return errors.New("dial handler non-exist")
	} else if n.listenHandler == nil && n.dialConn == nil {
		return errors.New("the obj is null")
	}
	return nil
}

func (n *NetChannel) Close() error {
	logger := log.New(constValue.logOutFile, "n.Close "+strconv.Itoa(n.port)+" ", log.Ldate|log.Ltime|log.Lshortfile)
	n.Wait()
	if n.listenHandler != nil {
		if err := n.listenHandler.Close(); err != nil {
			logger.Println(err)
			return err
		}
		if err := resourceLock.UnlockKey(n.ip + ":" + strconv.Itoa(n.port)); err != nil {
			logger.Println(err)
			return err
		}
	} else if n.dialConn != nil {
		if err := n.dialConn.Close(); err != nil {
			logger.Println(err)
			return err
		}
	}
	return nil
}

func (n *NetChannel) rsaEncrypt(publicKey *rsa.PublicKey, msg []byte) ([]byte, error) {
	hash := sha256.New()
	msgLen := len(msg)
	step := n.targetPublicKey.Size() - 2*hash.Size() - 2
	var encryptedBytes []byte
	var err error
	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		encryptedBlockBytes, err := rsa.EncryptOAEP(
			hash,
			rand.Reader,
			publicKey,
			msg[start:finish],
			nil)
		if err != nil {
			return encryptedBytes, err
		}
		encryptedBytes = append(encryptedBytes, encryptedBlockBytes...)
	}
	return encryptedBytes, err
}

func (n *NetChannel) rsaDecrypt(privatekey rsa.PrivateKey, encMsg []byte) ([]byte, error) {
	msgLen := len(encMsg)
	step := n.myPrivateKey.PublicKey.Size()
	var decryptedBytes []byte
	var err error
	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		decryptedBlockBytes, err := privatekey.Decrypt(
			nil,
			encMsg[start:finish],
			&rsa.OAEPOptions{Hash: crypto.SHA256})
		if err != nil {
			return decryptedBytes, err
		}

		decryptedBytes = append(decryptedBytes, decryptedBlockBytes...)
	}

	return decryptedBytes, err
}

func (n *NetChannel) rsaKeySetup(RsaKeyJson []byte, isPrivateKey bool) (*rsa.PrivateKey, error) {
	key := make(map[string]string)
	err := json.Unmarshal(RsaKeyJson, &key)
	if err != nil {
		return nil, err
	}

	PrivateKey := new(rsa.PrivateKey)
	PrivateKey.N = new(big.Int)
	tmp, err := base64.StdEncoding.DecodeString(key["N"])
	if err != nil {
		return nil, err
	}
	PrivateKey.N.GobDecode(tmp)
	tmp, err = base64.StdEncoding.DecodeString(key["E"])
	if err != nil {
		return nil, err
	}
	PrivateKey.E = int(binary.LittleEndian.Uint32(tmp))
	if isPrivateKey {
		PrivateKey.D = new(big.Int)
		tmp, err = base64.StdEncoding.DecodeString(key["D"])
		if err != nil {
			return nil, err
		}
		PrivateKey.D.GobDecode(tmp)
		primes := make([]*big.Int, 2)
		tmp := new(big.Int)
		tmp_prime, err := base64.StdEncoding.DecodeString(key["Primes1"])
		if err != nil {
			return nil, err
		}
		tmp.GobDecode(tmp_prime)
		primes[0] = tmp
		tmp = new(big.Int)
		tmp_prime, err = base64.StdEncoding.DecodeString(key["Primes2"])
		if err != nil {
			return nil, err
		}
		tmp.GobDecode(tmp_prime)
		primes[1] = tmp
		PrivateKey.Primes = primes
		PrivateKey.Precompute()
	}
	return PrivateKey, err
}

func GenRsaKeyPair() ([]byte, error) {
	logger := log.New(constValue.logOutFile, "GenRsaKeyPair ", log.Ldate|log.Ltime|log.Lshortfile)
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logger.Panic(err)
	}
	key := make(map[string]string)
	key["N"] = base64.StdEncoding.EncodeToString(func(n []byte, _ error) []byte { return n }(privateKey.N.GobEncode()))
	key["D"] = base64.StdEncoding.EncodeToString(func(n []byte, _ error) []byte { return n }(privateKey.D.GobEncode()))
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(privateKey.E))
	key["E"] = base64.StdEncoding.EncodeToString(bs)
	key["Primes1"] = base64.StdEncoding.EncodeToString(func(n []byte, _ error) []byte { return n }(privateKey.Primes[0].GobEncode()))
	key["Primes2"] = base64.StdEncoding.EncodeToString(func(n []byte, _ error) []byte { return n }(privateKey.Primes[1].GobEncode()))
	return json.Marshal(key)
}
