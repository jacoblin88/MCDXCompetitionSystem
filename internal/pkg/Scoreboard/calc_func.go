package Scoreboard

import "log"

func (sc *Scoreboard) CalcDefendScore() {

	sc.calcServiceDefScore(sc.ST.VPNSpyRecord_Arr)
	sc.calcServiceDefScore(sc.ST.ServiceCheckerRecord_Arr)
	log.Println("TeamScoreSet Now:")

	//sc.calcServiceDefScore(sc.ST.ServicePatchRecord_Arr)
	sc.calcFlagDefScore(sc.ST.FlagLostRecord_Arr)

	log.Print("Round Score summary:")
	log.Println(sc.ST.TeamScore)
}

func (sc *Scoreboard) calcFlagDefScore(record_arr []FlagRecord) {
	var machineid int
	var event_name string

	for _, s := range record_arr {
		machineid = s.Machineid

		for _, v := range sc.MachineInfo {
			if v.MachineID == machineid {
				event_name = s.EventName

				def_tmp := sc.ST.TeamScore[s.Teamid].Def
				atk_tmp := sc.ST.TeamScore[s.Teamid].Atk
				def_tmp += v.Event[event_name].Def

				//sc.ST.TeamScore[s.Teamid].Def = sc.ST.TeamScore[s.Teamid].Def + v.Event[event_name].Def
				sc.ST.TeamScore[s.Teamid] = TeamScoreSet{atk_tmp, def_tmp}
			}
		}
	}
}

func (sc *Scoreboard) calcServiceDefScore(record_arr []ServiceRecord) {

	var serviceid int
	var event_name string

	for _, s := range record_arr {
		//serviceid = sc.st.ServiceBlockRecord_Arr[i].Serviceid
		serviceid = s.Serviceid

		for _, v := range sc.ServiceInfo {
			if v.ServiceID == serviceid {

				//event_name = sc.st.ServiceBlockRecord_Arr[i].EventName

				event_name = s.EventName

				def_tmp := sc.ST.TeamScore[s.Teamid].Def
				atk_tmp := sc.ST.TeamScore[s.Teamid].Atk

				if event_name != "serviceblock" {
					def_tmp += int(v.Event[event_name].Def / 120)
				} else {
					def_tmp += int(v.Event[event_name].Def / 40)
				}
				//log.Println(sc.ST.TeamScore)
				//log.Println(TeamScoreSet{atk_tmp, def_tmp})
				sc.ST.TeamScore[s.Teamid] = TeamScoreSet{atk_tmp, def_tmp}
			}
		}
	}
}
