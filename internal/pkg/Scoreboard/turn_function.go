package Scoreboard

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/Config"
	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagOperator"
	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/SecretKey"
	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/ServiceChecker"
	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/VpnSpyMoniter"
)

var VpnSpyList = []struct {
	IP   string
	Port int
}{
	{
		IP:   "10.160.0.3",
		Port: 8443,
	},
	{
		IP:   "10.160.0.4",
		Port: 8443,
	},
	{
		IP:   "10.160.0.5",
		Port: 8443,
	},
	{
		IP:   "10.160.0.6",
		Port: 8443,
	},
	// {
	// 	IP:   "127.0.0.1",
	// 	Port: 8444,
	// },
	// {
	// 	IP:   "127.0.0.1",
	// 	Port: 8445,
	// },
	// {
	// 	IP:   "127.0.0.1",
	// 	Port: 8446,
	// },
}

func (sc *Scoreboard) srvcheck_operation() {

	var srv_chk *ServiceChecker.ServiceChecker
	for _, v := range sc.ServiceInfo {
		srv_chk = ServiceChecker.New(v.ServiceID, v.CheckNormal, v.CheckPatch, v.IP, v.Port, sc.Round, sc.Turn, sc.Servicecheck_chan)
		go srv_chk.PublishServiceState()
	}
}

func (sc *Scoreboard) vpnspy_operation() {
	var service_info []Config.ServiceInfoItem

	for _, v := range sc.ServiceInfo {
		var event = make(map[string]Config.EventItem)
		for name, item := range v.Event {
			event[name] = Config.EventItem{
				Atk: item.Atk,
				Def: item.Def,
			}
		}
		service_info = append(service_info, Config.ServiceInfoItem{
			ServiceID:   v.ServiceID,
			TeamID:      v.TeamID,
			MachineID:   v.MachineID,
			IP:          v.IP,
			Port:        v.Port,
			CheckPatch:  v.CheckPatch,
			CheckNormal: v.CheckNormal,
			Event:       event})
	}
	for i := 0; i < sc.ST.TeamNum; i++ {

		v, err := VpnSpyMoniter.New(VpnSpyList[i].IP, VpnSpyList[i].Port, i+1, service_info, sc.Round, sc.Turn, &sc.Vpnspy_chan, 10, 3, SecretKey.RsaSpyKey.PrivateKeyJsonStr, SecretKey.RsaSpyKey.PrivateKeyJsonStr)
		if err != nil {
			log.Println(err)
		}
		if err := v.Check(0); err != nil {
			log.Println(err)
		}
	}

}

func (sc *Scoreboard) flagcheck_operation() {

	for i := 0; i < len(sc.MachineInfo); i++ {
		yml := trigger(sc.MachineInfo[i])
		n, err := FlagOperator.New(yml, sc.MachineInfo[i].Flag_string,
			sc.Round,
			sc.Turn,
			&sc.Flag_chan,
			10, 3, SecretKey.RsaKey[0].PrivateKeyJsonStr,
			SecretKey.RsaKey[sc.MachineInfo[i].MachineID].PublicKeyJsonStr)
		if err != nil {
			log.Println(err)
		}

		go n.Check()
	}
}

func trigger(oldyml MachineInfoItem) Config.MachineInfoItem {
	var tmp = Config.FlagItem{
		Path: oldyml.Flag.Path,
		Port: oldyml.Flag.Port,
	}
	var ymlConfig = Config.MachineInfoItem{
		IP:        oldyml.IP,
		MachineID: oldyml.MachineID,
		TeamID:    oldyml.TeamID,
		Event:     nil,
		Flag:      tmp,
	}
	return ymlConfig
}
