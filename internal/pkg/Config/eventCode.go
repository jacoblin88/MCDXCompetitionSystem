package Config

type FlagItem struct {
	Path string `yaml:"path"`
	Port int    `yaml:"port"`
}

type EventItem struct {
	Atk int `yaml:"atk"`
	Def int `yaml:"def"`
}

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

type MachineInfoItem struct {
	MachineID int                  `yaml:"machineid"`
	TeamID    int                  `yaml:"teamid"`
	IP        string               `yaml:"ip"`
	Event     map[string]EventItem `yaml:"event,flow"`
	Flag      FlagItem             `yaml:"flag"`
	//flagloss;flagsubmit;resetsys
}
