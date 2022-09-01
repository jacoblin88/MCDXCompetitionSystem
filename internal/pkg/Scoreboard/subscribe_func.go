package Scoreboard

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/PangolinoLab/MCDXCompetitionSystem/pkg/DBlib"
	"github.com/PangolinoLab/MCDXCompetitionSystem/pkg/GameTimer"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func (sc *Scoreboard) SubscribeNewRound(pubSub gochannel.GoChannel, turn_gt GameTimer.GameTimer) {

	messages, _ := pubSub.Subscribe(context.Background(), "NewRound")

	go sc.processNewRound(messages, turn_gt)
	//return pubSub
}

func (sc *Scoreboard) processNewRound(messages <-chan *message.Message, turn_gt GameTimer.GameTimer) {

	for msg := range messages {
		log.Printf("Event:" + string(msg.Payload))

		//Turn off Start New Turn
		turn_gt.EventNameChan <- "StartNewRoundTurn"

		time.Sleep(3 * time.Second)

		db := DBlib.NewMySQL("configs/mysql/config.yaml")
		db_op := DBlib.DataBaseOperator(db)
		db_mysql := DBlib.SqlDatabase(db)

		//Calculate Score and insert into db
		sc.RoundCaculation()
		sc.InsertScoreRecordToDB(db_op, db_mysql)

		sc.Round += 1
		sc.Turn = 1

		//New flag Generate
		sc.GenerateNewFlagCollection()
		sc.UpdateAndInsertFlagToDB(db_op, db_mysql)

		sc.Flagupdate_operation()

		//Store log into DB
		sc.StoreLogToDB(db_op, db_mysql)

		//Initialize New Score Table
		//Start New Round
		sc.ST = *NewScoreTable(sc.ST.TeamNum)

		//Start up New Turn After dealing with round infomation
		turn_gt.EventNameChan <- "NewTurn"

		msg.Ack()
	}
}

func (sc *Scoreboard) SubscribeNewTurn(pubSub gochannel.GoChannel) {

	messages, _ := pubSub.Subscribe(context.Background(), "NewTurn")
	round_msg, _ := pubSub.Subscribe(context.Background(), "StartNewRoundTurn")

	go sc.processNewTurn(messages)
	go sc.processNewRoundTurn(round_msg)
}

func (sc *Scoreboard) processNewRoundTurn(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("%s,%s", msg.UUID, string(msg.Payload))
	}
}

func (sc *Scoreboard) processNewTurn(messages <-chan *message.Message) {

	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		//start Service check
		//sc.srvcheck_operation()

		//start vpnspy check
		//sc.vpnspy_operation()

		//start Flag check
		sc.flagcheck_operation()

		sc.Turn += 1

		msg.Ack()
	}
}

func (sc *Scoreboard) SubscribeVpnSpyMonitor() {
	//subscrbie event from VpnSpyMonitor Object

	messages, _ := sc.Vpnspy_chan.Subscribe(context.Background(), "VpnSpyMonitor")

	go sc.processVpnSpyMonitor(messages)
}

func (sc *Scoreboard) processVpnSpyMonitor(messages <-chan *message.Message) {

	for msg := range messages {

		var sb_record ServiceStatusRecord

		err := json.Unmarshal([]byte(msg.Payload), &sb_record)
		log.Print("Unmarshal VpnSpyMonitor Record:")
		log.Println(sb_record)

		if err != nil {
			fmt.Println("unmarshal:" + string(err.Error()))
			msg.Ack()
			continue
		}

		//Add record to scoreboard
		s_record := sc.TranslateVpnSpyRecord(sb_record)
		log.Print("Translate Result:")
		log.Println(s_record)
		sc.ST.VPNSpyRecord_Arr = append(sc.ST.VPNSpyRecord_Arr, s_record)

		//Update VpnSpy monitoring status
		sc.ServiceStatus.VpnSpy_Result[sb_record.ServiceID] = sb_record.Event

		msg.Ack()
	}
}

func (sc *Scoreboard) SubscribeServiceChecker() {
	//subscrbie event from VpnSpyMonitor Object

	messages, _ := sc.Servicecheck_chan.Subscribe(context.Background(), "ServiceChecker")

	go sc.processServiceChecker(messages)
}

func (sc *Scoreboard) processServiceChecker(messages <-chan *message.Message) {

	for msg := range messages {

		var sb_record ServiceStatusRecord
		json.Unmarshal([]byte(msg.Payload), &sb_record)

		//Add record to scoreboard
		s_record := sc.TranslateServiceCheckerRecord(sb_record)

		sc.ST.ServiceCheckerRecord_Arr = append(sc.ST.ServiceCheckerRecord_Arr, s_record)

		//Update ServiceChecker monitoring status
		sc.ServiceStatus.ServiceChecker_Result[sb_record.ServiceID] = sb_record.Event

		msg.Ack()
	}
}

func (sc *Scoreboard) SubscribeFlagCheckState() {
	//subscrbie event from VpnSpyMonitor Object

	messages, _ := sc.Flag_chan.Subscribe(context.Background(), "flagcheck")

	go sc.processFlagCheckState(messages)
}

func (sc *Scoreboard) processFlagCheckState(messages <-chan *message.Message) {

	for msg := range messages {
		log.Println(msg.Payload)

		var fck_record FlagStatusRecord
		json.Unmarshal([]byte(msg.Payload), &fck_record)

		log.Println(fck_record)
		//Add record to scoreboard
		f_record := sc.TranslateFlagCheckRecord(fck_record)
		sc.ST.FlagLostRecord_Arr = append(sc.ST.FlagLostRecord_Arr, f_record)

		//Update FlagCheckStatus
		sc.FlagStatus.FlagCheckStatus[fck_record.MachineID] = fck_record.Event

		msg.Ack()
	}
}

func (sc *Scoreboard) SubscribeFlagUpdateState() {
	//subscrbie event from VpnSpyMonitor Object

	messages, _ := sc.Flag_chan.Subscribe(context.Background(), "flagupdate")

	go sc.processFlagUpdateState(messages)
}

func (sc *Scoreboard) processFlagUpdateState(messages <-chan *message.Message) {

	for msg := range messages {

		var fck_record FlagStatusRecord
		json.Unmarshal([]byte(msg.Payload), &fck_record)

		//Add record to scoreboard
		f_record := sc.TranslateFlagStatusRecord(fck_record)
		sc.ST.FlagLostRecord_Arr = append(sc.ST.FlagLostRecord_Arr, f_record)

		//Update FlagUpdateStatus
		sc.FlagStatus.FlagUpdateStatus[fck_record.MachineID] = fck_record.Event

		msg.Ack()
	}
}
