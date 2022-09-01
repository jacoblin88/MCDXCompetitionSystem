package FlagOperator

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/Config"
	"github.com/PangolinoLab/MCDXCompetitionSystem/pkg/NetChannel"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

var constValue = struct {
	normal                int
	unNormal              int
	updateWatermilEventId string
	checkWatermilEventId  string
	actionCheck           string
	actionUpdate          string
	logOutFile            io.Writer
	eventLoggerOutFile    io.Writer
}{
	normal:                1,
	unNormal:              0,
	updateWatermilEventId: "flagupdate",
	checkWatermilEventId:  "flagcheck",
	actionCheck:           "check",
	actionUpdate:          "update",
	logOutFile:            os.Stdout,
	eventLoggerOutFile:    os.Stdout,
}

type FlagOperator struct {
	machineId     int
	machineInfo   Config.MachineInfoItem
	flag          string
	round         int
	turn          int
	pubSub        *gochannel.GoChannel
	logger        *log.Logger
	eventLogger   *log.Logger
	clientChannel *NetChannel.NetChannel
}

func New(machineInfo Config.MachineInfoItem, flag string, round int, turn int, pubSub *gochannel.GoChannel, timeout float32, tryTimes int, PrivateKeyConf string, PublicKeyConf string) (*FlagOperator, error) {
	// basic init
	var err error
	var f FlagOperator
	f.machineId = machineInfo.MachineID
	f.machineInfo = machineInfo
	f.flag = flag
	f.round = round
	f.turn = turn
	f.pubSub = pubSub
	f.logger = log.New(constValue.logOutFile, "FlagOperator "+strconv.Itoa(f.machineId)+" ", log.Ldate|log.Ltime|log.Lshortfile)
	f.eventLogger = log.New(constValue.eventLoggerOutFile, strconv.Itoa(f.machineId)+" ", log.Ldate|log.Ltime)
	// networking init
	f.clientChannel, err = NetChannel.New(f.machineInfo.IP, f.machineInfo.Flag.Port, timeout, tryTimes, PrivateKeyConf, PublicKeyConf)
	if err != nil {
		f.logger.Println(err)
		return nil, err
	}
	return &f, err
}

func (f *FlagOperator) Check() error {
	return f._impFlagOperator(constValue.actionCheck, f.recvCallBackCheck)
}

func (f *FlagOperator) Update() error {
	return f._impFlagOperator(constValue.actionUpdate, f.recvCallBackUpdate)
}

func (f *FlagOperator) _impFlagOperator(action string, callback NetChannel.RecvCallBackFunc) error {
	var netMsg Config.NetMsgMachineStruct
	netMsg.Action = action
	netMsg.MachineId = f.machineId
	netMsg.Event = 0
	netMsg.Flag = f.flag
	netMsg.Path = f.machineInfo.Flag.Path
	netMsg.Log = ""
	netMsg.Round = f.round
	netMsg.Turn = f.turn
	netMsgJson, err := json.Marshal(netMsg)
	if err != nil {
		f.logger.Println(err)
		return err
	}
	_, err = f.clientChannel.Send(netMsgJson, len(netMsgJson))
	if err != nil {
		f.eventLogger.Println("Send ", err)
		f.logger.Println(err)
		return err
	}
	f.eventLogger.Println("Send ", netMsgJson)
	_, err = f.clientChannel.Recv(callback)
	if err != nil {
		f.eventLogger.Println("Recv ", err)
		f.logger.Println(err)
		netMsg.Event = constValue.unNormal
		netMsg.Log = "target not respond"
		netMsgJson, err = json.Marshal(netMsg)
		if err != nil {
			f.eventLogger.Println(err)
			f.logger.Println(err)
			return err
		} else if netMsg.Action == constValue.actionCheck {
			f._recvCallBack(netMsgJson, len(netMsgJson), constValue.checkWatermilEventId)
		} else if netMsg.Action == constValue.actionUpdate {
			f._recvCallBack(netMsgJson, len(netMsgJson), constValue.updateWatermilEventId)
		}
		return err
	}
	if err := f.clientChannel.Close(); err != nil {
		f.logger.Println(err)
	}
	return err
}

func (f *FlagOperator) recvCallBackCheck(msg []byte, msgSz int) (respMsg []byte, err error) {
	return f._recvCallBack(msg, msgSz, constValue.checkWatermilEventId)
}

func (f *FlagOperator) recvCallBackUpdate(msg []byte, msgSz int) (respMsg []byte, err error) {
	return f._recvCallBack(msg, msgSz, constValue.updateWatermilEventId)
}

func (f *FlagOperator) _recvCallBack(msg []byte, msgSz int, action string) (respMsg []byte, err error) {
	f.eventLogger.Println("Recv ", string(msg))
	var resp Config.NetMsgMachineStruct
	if err := json.Unmarshal(msg, &resp); err != nil {
		f.logger.Println(err)
		return nil, err
	}
	jsonEvent, err := json.Marshal(resp.EventDrivenMachineStruct)
	if err != nil {
		f.logger.Println(err)
		return nil, err
	}
	eventMsg := message.NewMessage(watermill.NewUUID(), jsonEvent)

	if err := f.pubSub.Publish(action, eventMsg); err != nil {
		f.logger.Println(err)
		return nil, err
	}
	return nil, nil
}

//
// check <-> client: {"action": ["check"|"update"], 'machineid':int,'event':[0|1], 'flag': string, 'path': string, 'log': string, 'round':int, 'turn':int}
//
// watermil: {'machineid':int,'event':[0|1], 'round':int, 'turn':int}
