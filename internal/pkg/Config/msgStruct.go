package Config

type NetMsgMachineStruct struct {
	EventDrivenMachineStruct
	Action string
	Flag   string
	Path   string
	Log    string
}

type EventDrivenMachineStruct struct {
	MachineId int
	Event     int
	Round     int
	Turn      int
}

type NetMsgServiceStruct struct {
	EventDrivenServiceStruct
	IP   string
	Port int
	Path string
	Log  string
}

type EventDrivenServiceStruct struct {
	ServiceId int
	Event     int
	Round     int
	Turn      int
}
