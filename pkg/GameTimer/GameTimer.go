package GameTimer

import (
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type GameTimer struct {
	eventTime int
	EventName string
	Pubsub    *gochannel.GoChannel
	tick      *time.Ticker
	//**
	EventNameChan chan string
	//**
}

func NewGameTimer(sec int, eventName string) *GameTimer {
	gt := new(GameTimer)
	gt.eventTime = sec
	gt.EventName = eventName
	gt.Pubsub = gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)
	//**
	gt.EventNameChan = make(chan string, 1)
	//**
	return gt
}

func (gt *GameTimer) PublishTime() {
	//start publish time event
	gt.tick = time.NewTicker(time.Duration(gt.eventTime) * time.Second)
	go gt.publishTimeEvent_gorutine()
}

func (gt *GameTimer) publishTimeEvent_gorutine() {
	//**
	var EventName string = gt.EventName
	//**
	for {
		//**
		select {
		case <-time.After(time.Second * time.Duration(gt.eventTime)):
			if EventName != "" {
				msg := message.NewMessage(watermill.NewUUID(), []byte(EventName))

				if err := gt.Pubsub.Publish(gt.EventName, msg); err != nil {
					panic(err)
				}
			}
		case EventName = <-gt.EventNameChan:
		}

		//**
		// <-gt.tick.C
		// msg := message.NewMessage(watermill.NewUUID(), []byte(gt.EventName))

		// if err := gt.Pubsub.Publish(gt.EventName, msg); err != nil {
		// 	panic(err)
		// }
	}
}
