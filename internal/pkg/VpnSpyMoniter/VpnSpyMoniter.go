package VpnSpyMoniter

import (
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/Config"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
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
	eventLoggerOutFile:        ioutil.Discard,
}

type VpnSpyMoniter struct {
	vpnSpyIp       string
	vpnSpyPort     int
	teamId         int
	serviceInfos   []Config.ServiceInfoItem
	round          int
	turn           int
	timeout        float32
	tryTimes       int
	PrivateKeyConf string
	PublicKeyConf  string
	logger         *log.Logger
	eventLogger    *log.Logger
	pubSub         *gochannel.GoChannel
	workers        []*VpnSpyMoniterMicroWorker
}

func New(vpnSpyIp string, vpnSpyPort int, teamid int, serviceInfos []Config.ServiceInfoItem, round int, turn int, pubSub *gochannel.GoChannel, timeout float32, tryTimes int, PrivateKeyConf string, PublicKeyConf string) (*VpnSpyMoniter, error) {
	var v VpnSpyMoniter
	v.logger = log.New(constValue.logOutFile, "VpnSpyMoniter "+v.vpnSpyIp+":"+strconv.Itoa(v.vpnSpyPort)+" ", log.Ldate|log.Ltime|log.Lshortfile)
	v.eventLogger = log.New(constValue.eventLoggerOutFile, "VpnSpyMoniter "+v.vpnSpyIp+":"+strconv.Itoa(v.vpnSpyPort)+" ", log.Ldate|log.Ltime)
	v.serviceInfos = serviceInfos
	v.vpnSpyIp = vpnSpyIp
	v.vpnSpyPort = vpnSpyPort
	v.teamId = teamid
	v.round = round
	v.turn = turn
	v.pubSub = pubSub
	v.timeout = timeout
	v.tryTimes = tryTimes
	v.PrivateKeyConf = PrivateKeyConf
	v.PublicKeyConf = PublicKeyConf

	// TODO: in loop  for each service check method
	for _, serviceInfo := range v.serviceInfos {
		if serviceInfo.TeamID == v.teamId {
			continue
		}
		var netMsg = Config.NetMsgServiceStruct{
			IP:   serviceInfo.IP,
			Port: serviceInfo.Port,
			Path: serviceInfo.CheckNormal,
			Log:  "",
			EventDrivenServiceStruct: Config.EventDrivenServiceStruct{
				ServiceId: serviceInfo.ServiceID,
				Event:     0,
				Round:     v.round,
				Turn:      v.turn,
			},
		}
		var tmp *VpnSpyMoniterMicroWorker
		var err error
		tmp, err = _microWorkerNew(v.vpnSpyIp, v.vpnSpyPort, v.timeout, v.tryTimes, v.PrivateKeyConf, v.PublicKeyConf, netMsg, v.pubSub)
		if err != nil {
			v.logger.Println(err)
			v.eventLogger.Println(err)
			return nil, err
		}
		v.workers = append(v.workers, tmp)
	}
	return &v, nil
}

func (v *VpnSpyMoniter) Check(sleepTime float32) error {
	for _, worker := range v.workers {
		worker.Check()
		time.Sleep(time.Second * time.Duration(sleepTime))
	}
	return nil
}
