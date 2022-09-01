package VpnSpy

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/Config"
	"github.com/PangolinoLab/MCDXCompetitionSystem/pkg/NetChannel"
)

var constValue = struct {
	normal                    int
	unNormal                  int
	checkBlockWatermilEventId string
	logOutFile                io.Writer
	eventLoggerOutFile        io.Writer
}{
	normal:                    1,
	unNormal:                  0,
	checkBlockWatermilEventId: "VpnSpyMonitor",
	logOutFile:                ioutil.Discard,
	eventLoggerOutFile:        nil,
}

type VpnSpy struct {
	logger        *log.Logger
	eventLogger   *log.Logger
	ip            string
	port          int
	serverChannel *NetChannel.NetChannel
	quitChan      chan bool
}

func New(ip string, port int, tryTimes int, PrivateKeyConf string, PublicKeyConf string) (*VpnSpy, error) {
	var v VpnSpy
	v.ip = ip
	v.port = port
	v.quitChan = make(chan bool, 1)
	a, _ := os.Open("/var/log/VpnSpy/log.txt")
	constValue.eventLoggerOutFile = a
	v.logger = log.New(constValue.logOutFile, "VpnSpy "+v.ip+":"+strconv.Itoa(v.port)+" ", log.Ldate|log.Ltime|log.Lshortfile)
	v.eventLogger = log.New(constValue.eventLoggerOutFile, "VpnSpy "+v.ip+":"+strconv.Itoa(v.port)+" ", log.Ldate|log.Ltime)
	n, err := NetChannel.New(ip, port, -1, tryTimes, PrivateKeyConf, PublicKeyConf)
	if err != nil {
		v.logger.Panic(err)
		return nil, err
	}
	v.serverChannel = n
	return &v, err
}

func (v *VpnSpy) StartService() error {
	v.eventLogger.Println(errors.New("start service"))
	for {
		select {
		case <-v.quitChan:
			err := errors.New("recv quit signal")
			v.logger.Println(err)
			v.eventLogger.Println(err)
			return err
		default:
		}
		_, err := v.serverChannel.Recv(v.callBack)
		if err != nil {
			v.logger.Println(err)
			v.eventLogger.Println(err)
		}
	}
}

func (v *VpnSpy) StopService() error {
	v.quitChan <- true
	if err := v.serverChannel.Close(); err != nil {
		v.logger.Println(err)
		return err
	}
	v.eventLogger.Println(errors.New("stop service"))
	return nil
}

func (v *VpnSpy) callBack(msg []byte, msgSz int) ([]byte, error) {
	v.eventLogger.Println("Recv: " + string(msg[:msgSz]))
	var netMsg Config.NetMsgServiceStruct
	var err error
	if err = json.Unmarshal(msg[:msgSz], &netMsg); err != nil {
		v.logger.Println(err)
		return nil, err
	}
	isOk, err := checkBlock(netMsg.IP, netMsg.Port)
	if err != nil {
		netMsg.Event = constValue.normal
		netMsg.Log = err.Error()
	} else {
		switch isOk {
		case true:
			netMsg.Event = constValue.normal
		case false:
			netMsg.Event = constValue.unNormal
		}
		netMsg.Log = ""
	}

	respMsg, err := json.Marshal(netMsg)
	if err != nil {
		v.logger.Println(err)
		return nil, err
	}
	v.eventLogger.Println("Send: " + string(respMsg))
	return respMsg, nil
}
