package Scoreboard

import (
	"io/ioutil"
	"log"

	//"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/VpnSpyMonitor"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"gopkg.in/yaml.v2"
)

type Scoreboard struct {
	ServiceInfo []ServiceInfoItem
	MachineInfo []MachineInfoItem

	// for web services
	//map machineid to machine status
	//0:flagloss
	//1:flagexist
	FlagStatus    FlagStatus_struct
	ServiceStatus ServiceStatus_struct
	//FlagUpdateStatus map[int]int

	Flag_chan         gochannel.GoChannel
	Vpnspy_chan       gochannel.GoChannel
	Servicecheck_chan gochannel.GoChannel

	Round int
	Turn  int

	//Scoretable for every Round
	ST ScoreTable

	//MySQL Object for this scoreboard
}

type FlagStatus_struct struct {
	FlagCheckStatus  map[int]int
	FlagUpdateStatus map[int]int
}

type ServiceStatus_struct struct {
	VpnSpy_Result map[int]int
	//0:serviceblock
	//1:normal

	ServiceChecker_Result map[int]int
	//0:serviceblock
	//1:normal
	//2:patch
}

func NewScoreboard(sys_config string, machine_config string, teamnum int) *Scoreboard {

	p := new(Scoreboard)
	p.Round = 2
	p.Turn = 1

	//Initialize Service and MachineInfo
	sys_data, _ := ioutil.ReadFile(string(sys_config))
	err := yaml.Unmarshal([]byte(sys_data), &(p.ServiceInfo))

	if err != nil {
		log.Fatalf("Fail to Read serviceinfo.yaml: %v", err)
	}

	machine_data, _ := ioutil.ReadFile(string(machine_config))
	err = yaml.Unmarshal([]byte(machine_data), &(p.MachineInfo))

	if err != nil {
		log.Fatalf("Fail to Read machineinfo.yaml: %v", err)
	}

	//generate flag collection
	p.GenerateNewFlagCollection()

	//Initialize Service and Machine Status For Web Service
	p.ServiceStatus.ServiceChecker_Result = make(map[int]int)
	p.ServiceStatus.VpnSpy_Result = make(map[int]int)

	for _, v := range p.ServiceInfo {
		p.ServiceStatus.ServiceChecker_Result[v.ServiceID] = 1
		p.ServiceStatus.VpnSpy_Result[v.ServiceID] = 1
	}

	p.FlagStatus.FlagCheckStatus = make(map[int]int)
	p.FlagStatus.FlagUpdateStatus = map[int]int{}

	for _, v := range p.MachineInfo {
		p.FlagStatus.FlagCheckStatus[v.MachineID] = 1
		p.FlagStatus.FlagUpdateStatus[v.MachineID] = 1
	}

	//new all event channel
	p.Flag_chan = *gochannel.NewGoChannel(gochannel.Config{},
		watermill.NewStdLogger(false, false))
	p.Vpnspy_chan = *gochannel.NewGoChannel(gochannel.Config{},
		watermill.NewStdLogger(false, false))
	p.Servicecheck_chan = *gochannel.NewGoChannel(gochannel.Config{},
		watermill.NewStdLogger(false, false))

	p.ST = *NewScoreTable(teamnum)

	return p
}
