package FlagAgent

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/Config"
	"github.com/PangolinoLab/MCDXCompetitionSystem/pkg/NetChannel"
)

var constValue = struct {
	normal       int
	unNormal     int
	actionCheck  string
	actionUpdate string
	logOutFile   io.Writer
	windowsOs    string
	linuxOs      string
	macOs        string
}{
	normal:       1,
	unNormal:     0,
	actionCheck:  "check",
	actionUpdate: "update",
	logOutFile:   os.Stdout,
	windowsOs:    "windows",
	linuxOs:      "linux",
	macOs:        "darwin",
}

type FlagAgent struct {
	logger        *log.Logger
	ip            string
	port          int
	serverChannel *NetChannel.NetChannel
	quitChan      chan bool
	updateRound   int
	osPlatform    string
}

func New(ip string, port int, tryTimes int, PrivateKeyConf string, PublicKeyConf string) (*FlagAgent, error) {
	var f FlagAgent
	f.ip = ip
	f.port = port
	f.osPlatform = runtime.GOOS
	f.updateRound = 0
	f.quitChan = make(chan bool, 1)
	f.logger = log.New(constValue.logOutFile, "FlagOperator "+f.ip+":"+strconv.Itoa(f.port)+" ", log.Ldate|log.Ltime|log.Lshortfile)
	n, err := NetChannel.New(ip, port, -1, tryTimes, PrivateKeyConf, PublicKeyConf)
	if err != nil {
		f.logger.Panic(err)
		return nil, err
	}
	f.serverChannel = n
	return &f, err
}

func (f *FlagAgent) StartService() error {
	for {
		select {
		case <-f.quitChan:
			err := errors.New("recv quit signal")
			f.logger.Println(err)
			return err
		default:
		}
		_, err := f.serverChannel.Recv(f.recvCallBackRouter)
		if err != nil {
			f.logger.Println(err)
		}
	}
}

func (f *FlagAgent) StopService() error {
	f.quitChan <- true
	if err := f.serverChannel.Close(); err != nil {
		f.logger.Println(err)
		return err
	}
	return nil
}

func (f *FlagAgent) recvCallBackRouter(msg []byte, msgSz int) ([]byte, error) {
	var netMsg Config.NetMsgMachineStruct
	var err error
	if err = json.Unmarshal(msg[:msgSz], &netMsg); err != nil {
		return nil, err
	}
	if netMsg.Action == constValue.actionCheck {
		if f.updateRound <= netMsg.Round {
			f.updateRound = netMsg.Round
			netMsg.Event, err = f.checkFlagState(netMsg.Flag, netMsg.Path)
			netMsg.Log = err.Error()
		} else {
			err = errors.New("this check action is delayed")
			netMsg.Event = constValue.normal
			netMsg.Log = err.Error()
			f.logger.Println(err)
		}
	} else if netMsg.Action == constValue.actionUpdate {
		if f.updateRound <= netMsg.Round {
			netMsg.Event, err = f.updateFlag(netMsg.Flag, netMsg.Path)
			netMsg.Log = err.Error()
			f.updateRound = netMsg.Round
		} else {
			err = errors.New("already change the flag")
			netMsg.Event = constValue.normal
			netMsg.Log = err.Error()
			f.logger.Println(err)
		}
	} else {
		err = errors.New("unknow action code")
		f.logger.Println(err)
		return nil, err
	}
	respMsg, err := json.Marshal(netMsg)
	if err != nil {
		f.logger.Println(err)
		return nil, err
	}
	return respMsg, nil
}

func (f *FlagAgent) checkFlagState(flag string, path string) (int, error) {
	_, err := os.Stat(path)
	if err != nil {
		f.logger.Println("check flag status ", constValue.unNormal, " ", err)
		return constValue.unNormal, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		f.logger.Println("check flag status ", constValue.unNormal, " ", err)
		return constValue.unNormal, err
	}
	if string(data) != flag {
		err = errors.New("flag string is wrong(" + string(data) + ")")
		f.logger.Println("check flag status ", constValue.unNormal, " ", err)
		return constValue.unNormal, err
	}
	f.logger.Println("check flag status ", constValue.normal, errors.New("flag is right"))
	return constValue.normal, errors.New("check " + path + " is right(" + string(data) + ")")
}

func (f *FlagAgent) updateFlag(flag string, path string) (int, error) {
	if err := os.WriteFile(path, []byte(flag), 0640); err != nil {
		f.logger.Println("update flag status ", constValue.unNormal, " ", err)
		return constValue.unNormal, err
	}
	f.logger.Println("update flag status ", constValue.normal, errors.New("flag is changed"))
	return constValue.normal, errors.New("update " + path + " content(" + string(flag) + ")")
}
