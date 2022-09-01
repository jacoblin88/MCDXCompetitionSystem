package Scoreboard

import (
	"crypto/rand"
	"math/big"
)

func (sc *Scoreboard) GenerateNewFlagCollection() {
	const letters = "0123456789abcdefghijklmnopqrstuvwxyz"

	const md5_len = 32

	for i := 0; i < len(sc.MachineInfo); i++ {
		sc.MachineInfo[i].Flag_string = sc.generateRandomString(md5_len)
	}
}

func (sc *Scoreboard) generateRandomString(n int) string {
	const letters = "0123456789abcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret)
}

type LogRecord struct {
	TeamID    int
	MachineID int
	Event     string
	Round     int
	Turn      int
}

//================Start Service and machine info======================
type ServiceInfoItem struct {
	ServiceID   int                  `yaml:"serviceid"`
	TeamID      int                  `yaml:"teamid"`
	MachineID   int                  `yaml:"machineid"`
	IP          string               `yaml:"ip"`
	Port        int                  `yaml:"port"`
	CheckPatch  string               `yaml:"checkpatch"`
	CheckNormal string               `yaml:"checknormal"`
	Event       map[string]EventItem `yaml:"event,flow"`
	//serviceloss,servicenormal,servicepatch,servicenotpatch
	//serviceblock,servicenotblock
}
type FlagItem struct {
	Path string `yaml:"path"`
	Port int    `yaml:"port"`
}
type MachineInfoItem struct {
	MachineID   int                  `yaml:"machineid"`
	TeamID      int                  `yaml:"teamid"`
	IP          string               `yaml:"ip"`
	Flag        FlagItem             `yaml:"flag"`
	Event       map[string]EventItem `yaml:"event,flow"`
	Flag_string string
	//flagloss;flagsubmit;resetsys
}

type EventItem struct {
	Atk int `yaml:"atk"`
	Def int `yaml:"def"`
}

//================End Service and machine info======================

//=====================start Score Table struct====================

type ScoreTable struct {
	TeamNum   int
	TeamScore map[int]TeamScoreSet

	//Store FlagOperator object detecting result
	FlagLostRecord_Arr []FlagRecord

	//Store ServiceChecker object detetecting result
	//event normal
	ServiceCheckerRecord_Arr []ServiceRecord
	//return event_name:serviceloss,servicenormal

	//Store VpnSpy object detecting result
	VPNSpyRecord_Arr []ServiceRecord
	//retunr event_name:serviceblock,servicenotblock
}

func NewScoreTable(teamnum int) *ScoreTable {

	p := new(ScoreTable)
	p.TeamNum = teamnum

	p.TeamScore = make(map[int]TeamScoreSet)

	for i := 1; i <= p.TeamNum; i++ {
		p.TeamScore[i] = TeamScoreSet{0, 0}
	}

	p.FlagLostRecord_Arr = make([]FlagRecord, 0)

	p.ServiceCheckerRecord_Arr = make([]ServiceRecord, 0)

	//p.ServicePatchRecord_Arr = make([]ServiceRecord, 0)

	p.VPNSpyRecord_Arr = make([]ServiceRecord, 0)

	return p
}

type TeamScoreSet struct {
	Atk int
	Def int
}

type ServiceStatusRecord struct {
	ServiceID int
	Round     int
	Turn      int
	Event     int
}

//Struct for Calc Score
type ServiceRecord struct {
	Teamid    int
	Machineid int
	Serviceid int

	IP   string
	Port int

	Round int
	Turn  int

	EventName string
	//event_name:
	//serviceloss,servicenormal
	//servicepatch,servicenotpatch
	//serviceblock,servicenotblock

	//VpnSpyMonitor
	//output: serviceid,Round,Turn,EventName
	//others will get other info from serviceinfo_arr
}

func (sc *Scoreboard) TranslateVpnSpyRecord(sb_record ServiceStatusRecord) ServiceRecord {
	var s_record ServiceRecord
	s_record.Serviceid = sb_record.ServiceID
	s_record.Round = sb_record.Round
	s_record.Turn = sb_record.Turn

	for _, v := range sc.ServiceInfo {
		if v.ServiceID == sb_record.ServiceID {
			s_record.IP = v.IP
			s_record.Machineid = v.MachineID
			s_record.Port = v.Port
			s_record.Teamid = v.TeamID

			if sb_record.Event == 0 {
				s_record.EventName = "serviceblock"
			} else {
				s_record.EventName = "servicenotblock"
			}
			break
		}
	}
	return s_record
}

func (sc *Scoreboard) TranslateServiceCheckerRecord(sb_record ServiceStatusRecord) ServiceRecord {
	var s_record ServiceRecord
	s_record.Serviceid = sb_record.ServiceID
	s_record.Round = sb_record.Round
	s_record.Turn = sb_record.Turn

	for _, v := range sc.ServiceInfo {
		if v.ServiceID == sb_record.ServiceID {
			s_record.IP = v.IP
			s_record.Machineid = v.MachineID
			s_record.Port = v.Port
			s_record.Teamid = v.TeamID

			if sb_record.Event == 0 {
				s_record.EventName = "serviceloss"
			} else if sb_record.Event == 1 {
				s_record.EventName = "servicenormal"
			} else if sb_record.Event == 2 {
				s_record.EventName = "servicepatch"
			}
			break
		}
	}
	return s_record
}

type FlagStatusRecord struct {
	MachineID int
	Event     int
	Round     int
	Turn      int
}

// Struct for calculating score
type FlagRecord struct {
	Teamid    int
	Machineid int
	IP        string

	Round     int //20 min
	Turn      int //3s
	EventName string
	//flagloss,flagexist,flagsubmit,resetsys
}

func (sc *Scoreboard) TranslateFlagCheckRecord(fck_record FlagStatusRecord) FlagRecord {
	var f_record FlagRecord

	f_record.Machineid = fck_record.MachineID
	f_record.Round = fck_record.Round
	f_record.Turn = fck_record.Turn

	for _, v := range sc.MachineInfo {
		if v.MachineID == fck_record.MachineID {
			f_record.IP = v.IP
			f_record.Teamid = v.TeamID

			if fck_record.Event == 0 {
				f_record.EventName = "flagloss"
			} else if fck_record.Event == 1 {
				f_record.EventName = "flagexist"
			}
			break
		}
	}
	return f_record
}

func (sc *Scoreboard) TranslateFlagStatusRecord(fck_record FlagStatusRecord) FlagRecord {
	var f_record FlagRecord

	f_record.Machineid = fck_record.MachineID
	f_record.Round = fck_record.Round
	f_record.Turn = fck_record.Turn

	for _, v := range sc.MachineInfo {
		if v.MachineID == fck_record.MachineID {
			f_record.IP = v.IP
			f_record.Teamid = v.TeamID

			if fck_record.Event == 0 {
				f_record.EventName = "flagupdatefail"
			} else {
				f_record.EventName = "flagupdatesuc"
			}
			break
		}
	}
	return f_record
}

//======================end of ScoreTable struct==========

//===================start of Machine Condition===========
