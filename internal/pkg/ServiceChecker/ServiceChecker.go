package ServiceChecker

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type ServiceChecker struct {
	serviceid         int
	chk_normal_script string
	chk_patch_script  string
	normal_result     int
	patch_result      int
	ip                string
	port              int
	round             int
	turn              int
	chan_srv_chk      gochannel.GoChannel
}

type ServiceCheckRecord struct {
	ServiceID int
	Round     int
	Turn      int
	Event     int
}

func New(srv_id int, chk_normal_script string, chk_patch_script, ip string, port int, round int, turn int, chan_srv_chk gochannel.GoChannel) *ServiceChecker {
	p := new(ServiceChecker)
	p.serviceid = srv_id
	p.chk_normal_script = chk_normal_script
	p.chk_patch_script = chk_patch_script
	p.chan_srv_chk = chan_srv_chk
	p.round = round
	p.turn = turn
	return p
}

func (s_checker *ServiceChecker) PublishServiceState() {

	wg := new(sync.WaitGroup)

	wg.Add(2)
	go s_checker.startchecknormal(wg)
	go s_checker.startcheckpatch(wg)
	wg.Wait()

	s_checker.publishServiceCheckerEvent()
}

func (s_checker *ServiceChecker) publishServiceCheckerEvent() {

	var record ServiceCheckRecord
	record.ServiceID = s_checker.serviceid
	record.Round = s_checker.round
	record.Turn = s_checker.turn
	//log.Println("MyTurn:")
	//log.Println(s_checker.turn)

	if s_checker.normal_result == 0 {
		record.Event = 0
	} else if s_checker.normal_result == 1 && s_checker.patch_result == 1 {
		record.Event = 2
	} else {
		record.Event = 1
	}

	msg_str, err := json.Marshal(record)

	if err != nil {
		fmt.Println(err)
	}

	msg := message.NewMessage(watermill.NewUUID(), []byte(msg_str))

	if err := s_checker.chan_srv_chk.Publish("ServiceChecker", msg); err != nil {
		log.Println(err)
	}
}

func (s_checker *ServiceChecker) startchecknormal(wg *sync.WaitGroup) {
	defer wg.Done()

	//cmd := exec.Command("python3", s_checker.chk_normal_script, s_checker.ip, strconv.Itoa(s_checker.port))
	cmd := exec.Command(s_checker.chk_normal_script, s_checker.ip, strconv.Itoa(s_checker.port))

	fmt.Print(cmd.Path + cmd.Dir)

	if err := cmd.Start(); err != nil {
		log.Println("start check normal error: %v", err)
	}

	done := make(chan error)
	go func() { done <- cmd.Wait() }()
	select {
	case err := <-done:
		if err != nil {
			if exiterr, ok := err.(*exec.ExitError); ok {
				if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
					s_checker.normal_result = status.ExitStatus()
				}
			} else {
				log.Println("script no exit code: %v", err)
				s_checker.normal_result = 1
			}
		}
		// exited
	case <-time.After(10 * time.Second):
		s_checker.normal_result = 1
		// timed out
		// suppose it is some mistake on our program
		// => normal
	}
}

func (s_checker *ServiceChecker) startcheckpatch(wg *sync.WaitGroup) {
	defer wg.Done()

	cmd := exec.Command(s_checker.chk_patch_script, s_checker.ip, strconv.Itoa(s_checker.port))

	if err := cmd.Start(); err != nil {
		log.Println("start check patch error: %v", err)
	}

	done := make(chan error)

	go func() { done <- cmd.Wait() }()
	select {
	case err := <-done:
		if err != nil {
			if exiterr, ok := err.(*exec.ExitError); ok {
				if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
					s_checker.patch_result = status.ExitStatus()
					//log.Printf("Exit Status: %d", status.ExitStatus())
				}
			} else {
				log.Println("script no exit code: %v", err)
				s_checker.patch_result = 0
			}
		}
		// exited
	case <-time.After(10 * time.Second):
		s_checker.patch_result = 0
		// timed out
		//-> think it is not patch
	}
}
