package packets

type CommandResponsePacket struct {
	Content string `json:"content"`
	Command string `json:"command"`
}

type StatusResponsePacket struct {
	Request string `json:"request"`
	Status  string `json:"status"`
}

type TaskLaunchPacket struct {
	Request string `json:"request"`
	TaskId  string `json:"TaskId"`
}
