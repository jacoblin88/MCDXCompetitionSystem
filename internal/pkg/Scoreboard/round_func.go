package Scoreboard

import (
	"log"
	"strconv"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/FlagOperator"
	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/SecretKey"
	"github.com/PangolinoLab/MCDXCompetitionSystem/pkg/DBlib"
	"github.com/bradfitz/slice"
)

func (sc *Scoreboard) Flagupdate_operation() {

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
		go n.Update()
	}
}

func (sc *Scoreboard) RoundCaculation() {
	sc.CalcDefendScore()
	//sc.CalcAtkScore()
}

func (sc *Scoreboard) InsertScoreRecordToDB(db_op DBlib.DataBaseOperator,
	db_mysql DBlib.SqlDatabase) {
	err := db_mysql.CheckStatus()
	if err != nil {
		log.Print(err)
	} else {
		//into mysql record
		var teamid []int = make([]int, 0)
		var atkscore []int = make([]int, 0)
		var defscore []int = make([]int, 0)
		var round []int = make([]int, 0)

		for v := range sc.ST.TeamScore {
			teamid = append(teamid, v)
			atkscore = append(atkscore, sc.ST.TeamScore[v].Atk)
			defscore = append(defscore, sc.ST.TeamScore[v].Def)
			round = append(round, sc.Round)
		}

		err := db_op.InsertScoreRecord(teamid, atkscore, defscore, round)
		if err != nil {
			log.Println(err)
		}
	}
}

func (sc *Scoreboard) UpdateAndInsertFlagToDB(db_op DBlib.DataBaseOperator,
	db_mysql DBlib.SqlDatabase) {
	err := db_mysql.CheckStatus()

	if err != nil {
		log.Print(err)
	} else {
		//into mysql record
		var teamid []int = make([]int, 0)
		var machineid []int = make([]int, 0)
		var flag []string = make([]string, 0)
		var round []int = make([]int, 0)

		for _, machine := range sc.MachineInfo {
			teamid = append(teamid, machine.TeamID)
			machineid = append(machineid, machine.MachineID)
			flag = append(flag, machine.Flag_string)
			round = append(round, sc.Round)
		}

		//new flag into mysql

		err := db_op.UpdateFlagTable(teamid, machineid, flag, round)
		if err != nil {
			log.Println(err)
		}
		err = db_op.InsertToFlagRecord(teamid, machineid, flag, round)
		if err != nil {
			log.Println(err)
		}
	}
}

func (sc *Scoreboard) StoreLogToDB(db_op DBlib.DataBaseOperator,
	db_mysql DBlib.SqlDatabase) {

	//	EventType

	// 0:flag lost

	// serviceID:0
	// result:
	// 0:flagcheckloss
	// 1:flagchecknormal
	// 2:flagupdatesuccess
	// 3:flagupdtefail
	// 1:ServiceCheck:

	// result:
	// 0:normalcheckfail
	// 1:normalchecksuccess
	// 2:patchchecksuccess
	// 3:patchcheckfail
	// 2:VpnSpy:

	log_arr := make([]LogRecord, 0)
	err := db_mysql.CheckStatus()
	if err != nil {
		log.Print(err)
	} else {
		// //Insert FlagLost Record
		teamid := make([]int, 0)
		machineid := make([]int, 0)
		event := make([]string, 0)
		round := make([]int, 0)
		turn := make([]int, 0)

		for _, v := range sc.ST.FlagLostRecord_Arr {
			log_record := LogRecord{v.Teamid, v.Machineid, v.EventName + "-0", v.Round, v.Turn}
			log_arr = append(log_arr, log_record)
		}

		for _, v := range sc.ST.ServiceCheckerRecord_Arr {
			log_record := LogRecord{v.Teamid, v.Machineid, v.EventName + "-" + strconv.Itoa(v.Serviceid), v.Round, v.Turn}
			log_arr = append(log_arr, log_record)
		}

		for _, v := range sc.ST.VPNSpyRecord_Arr {
			log_record := LogRecord{v.Teamid, v.Machineid, v.EventName + "-" + strconv.Itoa(v.Serviceid), v.Round, v.Turn}
			log_arr = append(log_arr, log_record)
		}
		slice.Sort(log_arr[:], func(i, j int) bool {
			return log_arr[i].TeamID < log_arr[j].TeamID
		})

		teamid_tmp := 1
		for _, v := range log_arr {
			if v.TeamID > teamid_tmp {
				teamid_tmp = v.TeamID
				err := db_op.InsertToLogTable(teamid, machineid, event, round, turn)
				if err != nil {
					log.Println("fail to store log:" + string(err.Error()))
				}
				teamid = make([]int, 0)
				machineid = make([]int, 0)
				event = make([]string, 0)
				round = make([]int, 0)
				turn = make([]int, 0)
			} else {
				teamid = append(teamid, v.TeamID)
				machineid = append(machineid, v.MachineID)
				round = append(round, v.Round)
				turn = append(turn, v.Turn)
				event = append(event, v.Event)
			}
		}
		err := db_op.InsertToLogTable(teamid, machineid, event, round, turn)
		if err != nil {
			log.Println("fail to store log:" + string(err.Error()))
		}

	}

}
