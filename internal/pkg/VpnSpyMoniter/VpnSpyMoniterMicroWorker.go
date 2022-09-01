package VpnSpyMoniter

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/Config"
	"github.com/PangolinoLab/MCDXCompetitionSystem/pkg/NetChannel"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type VpnSpyMoniterMicroWorker struct {
	ip            string
	port          int
	timeout       float32
	tryTimes      int
	privateKey    string
	publicKey     string
	netMsg        Config.NetMsgServiceStruct
	pubSub        *gochannel.GoChannel
	clientChannel *NetChannel.NetChannel
	logger        *log.Logger
	eventLogger   *log.Logger
}

func _microWorkerNew(ip string, port int, timeout float32, tryTimes int, privateKey string, publicKey string, netMsg Config.NetMsgServiceStruct, pubSub *gochannel.GoChannel) (*VpnSpyMoniterMicroWorker, error) {
	var w VpnSpyMoniterMicroWorker
	w.logger = log.New(constValue.logOutFile, "_microWorkerNew "+strconv.Itoa(netMsg.ServiceId)+" ", log.Ldate|log.Ltime|log.Lshortfile)
	w.eventLogger = log.New(constValue.eventLoggerOutFile, "_microWorkerNew "+strconv.Itoa(netMsg.ServiceId)+" ", log.Ldate|log.Ltime)
	w.ip = ip
	w.port = port
	w.timeout = timeout
	w.tryTimes = tryTimes
	w.privateKey = privateKey
	w.publicKey = publicKey
	w.netMsg = netMsg
	w.pubSub = pubSub
	tmp, err := NetChannel.New(w.ip, w.port, w.timeout, w.tryTimes, w.privateKey, w.publicKey)
	if err != nil {
		w.logger.Println(err)
		w.eventLogger.Println(err)
		return nil, err
	}
	w.clientChannel = tmp
	return &w, err
}

func (w *VpnSpyMoniterMicroWorker) Check() error {
	netMsgJson, err := json.Marshal(w.netMsg)
	if err != nil {
		w.logger.Println(err)
		return err
	}
	_, err = w.clientChannel.Send(netMsgJson, len(netMsgJson))
	if err != nil {
		w.eventLogger.Println("Send ", err)
		w.logger.Println(err)
		return err
	}
	w.eventLogger.Println("Send ", netMsgJson)
	_, err = w.clientChannel.Recv(w.recvCallBack)
	if err != nil {
		w.eventLogger.Println("Recv ", err)
		w.logger.Println(err)
		w.netMsg.Event = constValue.normal
		w.netMsg.Log = "target not respond"
		netMsgJson, err = json.Marshal(w.netMsg)
		if err != nil {
			w.logger.Println(err)
			w.eventLogger.Println(err)
			return err
		}
		w.recvCallBack(netMsgJson, len(netMsgJson))
		return err
	}
	if err := w.clientChannel.Close(); err != nil {
		w.logger.Println(err)
	}
	return err
}

func (w *VpnSpyMoniterMicroWorker) recvCallBack(msg []byte, msgSz int) (respMsg []byte, err error) {
	w.eventLogger.Println("Recv ", string(msg))
	var resp Config.NetMsgServiceStruct
	if err := json.Unmarshal(msg, &resp); err != nil {
		w.logger.Println(err)
		return nil, err
	}
	jsonEvent, err := json.Marshal(resp.EventDrivenServiceStruct)
	if err != nil {
		w.logger.Println(err)
		return nil, err
	}
	eventMsg := message.NewMessage(watermill.NewUUID(), jsonEvent)
	if err := w.pubSub.Publish(constValue.checkBlockWatermilEventId, eventMsg); err != nil {
		w.logger.Println(err)
		return nil, err
	}
	return nil, nil
}
