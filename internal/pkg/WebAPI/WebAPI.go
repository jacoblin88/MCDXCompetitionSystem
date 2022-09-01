package Webapi

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PangolinoLab/MCDXCompetitionSystem/internal/pkg/Scoreboard"
	"github.com/gorilla/mux"
)

type WebAPI struct {
	bindIP        string
	bindPort      int
	flagstatus    Scoreboard.FlagStatus_struct
	servicestatus Scoreboard.ServiceStatus_struct
	scoreboard    Scoreboard.Scoreboard
	srv           http.Server
	r             *mux.Router
}

func New(bindIP string, bindPort int, flagstatus Scoreboard.FlagStatus_struct,
	servicestatus Scoreboard.ServiceStatus_struct,
	scoreboard Scoreboard.Scoreboard) *WebAPI {

	p := new(WebAPI)
	p.bindIP = bindIP
	p.bindPort = bindPort
	p.flagstatus = flagstatus
	p.servicestatus = servicestatus
	p.scoreboard = scoreboard

	p.r = mux.NewRouter()
	p.r.HandleFunc("/service", p.getServiceStatus).Methods("POST")
	p.r.HandleFunc("/machine", p.getMachineStatus).Methods("POST")

	p.srv = http.Server{
		Handler:      p.r,
		Addr:         p.bindIP + ":" + strconv.Itoa(p.bindPort),
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	return p
}

type Web_ServiceStatus struct {
	ServiceID int
	Status    int
}

type Web_MachineStatus struct {
	MachineID int
	Status    int
}

func (webapi WebAPI) StartListen() {
	err := webapi.srv.ListenAndServe()
	fmt.Println(err)
}

func (webapi WebAPI) StopListen() {
	webapi.srv.Close()
}

func (webapi WebAPI) strByXOR(message string) string {
	messageLen := len(message)
	keywords := "430EFA4490BE8621AD9F4F325F19495D"
	keywordsLen := len(keywords)

	var result string

	for i := 0; i < messageLen; i++ {
		result += string(message[i] ^ keywords[i%keywordsLen])
	}

	sEnc := b64.StdEncoding.EncodeToString([]byte(result))
	return sEnc
}

func (webapi WebAPI) getMachineStatus(w http.ResponseWriter, r *http.Request) {

	teamid_str := r.PostFormValue("id")
	record_arr := make([]Web_MachineStatus, 0)
	var record Web_MachineStatus

	if teamid, err := strconv.Atoi(teamid_str); err == nil && teamid != 0 {
		machine_id_arr := make([]int, 0)

		for _, v := range webapi.scoreboard.ServiceInfo {
			if v.TeamID == teamid {
				machine_id_arr = append(machine_id_arr, v.ServiceID)
			}
		}

		for _, machineid := range machine_id_arr {

			if _, ok := webapi.flagstatus.FlagCheckStatus[machineid]; ok {
				//fmt.Println(webapi.servicestatus.ServiceChecker_Result[id])
				flagcheck_result := webapi.flagstatus.FlagCheckStatus[machineid]
				//flagupdate_result := webapi.flagstatus.FlagUpdateStatus[id]

				//checker=patch and vpnspy=normal => patch
				if flagcheck_result == 1 {
					record = Web_MachineStatus{machineid, 1}
				} else {
					record = Web_MachineStatus{machineid, 0}
				}
				record_arr = append(record_arr, record)

			}

		}

	} else {

		for v, _ := range webapi.flagstatus.FlagCheckStatus {

			//fmt.Println(v)
			flagcheck_result := webapi.servicestatus.ServiceChecker_Result[v]
			//flag_update_result := webapi.servicestatus.VpnSpy_Result[v]

			if flagcheck_result == 1 {
				record = Web_MachineStatus{v, 1}
			} else {
				record = Web_MachineStatus{v, 0}
			}

			record_arr = append(record_arr, record)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	jsondata, _ := json.Marshal(record_arr)
	result := webapi.strByXOR(string(jsondata))
	w.Write([]byte(result))
}

func (webapi WebAPI) getServiceStatus(w http.ResponseWriter, r *http.Request) {

	teamid_str := r.PostFormValue("id")
	record_arr := make([]Web_ServiceStatus, 0)

	var record Web_ServiceStatus

	if teamid, err := strconv.Atoi(teamid_str); err == nil && teamid != 0 {
		service_id_arr := make([]int, 0)

		//get service of the team
		for _, v := range webapi.scoreboard.ServiceInfo {
			if v.TeamID == teamid {
				service_id_arr = append(service_id_arr, v.ServiceID)
			}
		}

		for _, service_id := range service_id_arr {
			if _, ok := webapi.servicestatus.ServiceChecker_Result[service_id]; ok {
				//fmt.Println(webapi.servicestatus.ServiceChecker_Result[id])

				checker_result := webapi.servicestatus.ServiceChecker_Result[service_id]
				vpnspy_result := webapi.servicestatus.VpnSpy_Result[service_id]

				//checker=patch and vpnspy=normal => patch
				if checker_result == 0 || vpnspy_result == 0 {
					record = Web_ServiceStatus{service_id, 0}

				} else if checker_result == 2 && vpnspy_result == 1 {
					record = Web_ServiceStatus{service_id, 2}
				} else if checker_result == 1 && vpnspy_result == 1 {
					record = Web_ServiceStatus{service_id, 1}
				}
				record_arr = append(record_arr, record)
			}

		}

	} else {

		for v, _ := range webapi.servicestatus.ServiceChecker_Result {
			//fmt.Println(v)
			checker_result := webapi.servicestatus.ServiceChecker_Result[v]
			vpn_result := webapi.servicestatus.VpnSpy_Result[v]

			if checker_result == 0 || vpn_result == 0 {
				record = Web_ServiceStatus{v, 0}
			} else if checker_result == 2 && vpn_result == 1 {
				record = Web_ServiceStatus{v, 2}
			} else if checker_result == 1 && vpn_result == 1 {
				record = Web_ServiceStatus{v, 1}
			}
			record_arr = append(record_arr, record)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	jsondata, _ := json.Marshal(record_arr)
	result := webapi.strByXOR(string(jsondata))
	w.Write([]byte(result))
}
