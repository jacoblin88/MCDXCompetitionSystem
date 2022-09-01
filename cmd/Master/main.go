package main

import (
	"log"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/Scoreboard"
	Webapi "github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/WebAPI"
	"github.com/PangolinoLab/MCDXCompetitionSystem/pkg/DBlib"
	"github.com/PangolinoLab/MCDXCompetitionSystem/pkg/GameTimer"
)

var Service_Config = "configs/serviceinfo.yaml"
var Machine_Config = "configs/machineinfo.yaml"
var Mysql_Config = "configs/mysql/config.yaml"
var WebAPI_IP = "0.0.0.0"
var WebAPI_Port = 6666

func main() {

	Round := 60 * 20
	Turn := 30

	master := New(4, Round, Turn)
	master.StartService()

}

type Master struct {
	scoreboard Scoreboard.Scoreboard
	teamnum    int
	Roundtime  int
	TurnTime   int
}

func New(teamnum int, Roundtime int, TurnTime int) *Master {
	p := new(Master)
	p.teamnum = teamnum
	p.Roundtime = Roundtime
	p.TurnTime = TurnTime

	s := Scoreboard.NewScoreboard(Service_Config, Machine_Config, teamnum)
	p.scoreboard = *s

	return p
}

func (master *Master) StartService() {

	master.scoreboard.SubscribeFlagCheckState()
	master.scoreboard.SubscribeFlagUpdateState()
	master.scoreboard.SubscribeServiceChecker()
	master.scoreboard.SubscribeVpnSpyMonitor()

	//initialize flag and update it to machine
	master.init_Flag()

	turn_gt := GameTimer.NewGameTimer(master.TurnTime, "NewTurn")
	round_gt := GameTimer.NewGameTimer(master.Roundtime, "NewRound")

	master.scoreboard.SubscribeNewRound(*round_gt.Pubsub, *turn_gt)
	master.scoreboard.SubscribeNewTurn(*turn_gt.Pubsub)

	turn_gt.PublishTime()
	round_gt.PublishTime()

	wapi := Webapi.New(WebAPI_IP, WebAPI_Port, master.scoreboard.FlagStatus, master.scoreboard.ServiceStatus, master.scoreboard)
	wapi.StartListen()
	// <-time.After(time.Second * 600)

}

func (master *Master) init_Flag() {
	var teamid []int = make([]int, 0)
	var machineid []int = make([]int, 0)
	var flag []string = make([]string, 0)
	var round []int = make([]int, 0)

	master.scoreboard.Flagupdate_operation()

	for _, machine := range master.scoreboard.MachineInfo {
		teamid = append(teamid, machine.TeamID)
		machineid = append(machineid, machine.MachineID)
		flag = append(flag, machine.Flag_string)
		round = append(round, master.scoreboard.Round)
	}

	db := DBlib.NewMySQL(Mysql_Config)
	db_op := DBlib.DataBaseOperator(db)
	db_mysql := DBlib.SqlDatabase(db)

	err := db_mysql.CheckStatus()

	if err != nil {
		log.Println(err)
	} else {
		err = db_op.InsertToFlagTable(teamid, machineid, flag, round)
	}

	if err != nil {
		log.Println(err)
		err = db_op.UpdateFlagTable(teamid, machineid, flag, round)
	}

	if err != nil {
		log.Println(err)
	}

	err = db_op.InsertToFlagRecord(teamid, machineid, flag, round)

	if err != nil {
		log.Println(err)
	}

}
