package packets

const (
	TypePlayVideoRequestPacket   = "PlayVideoRequest"
	TypePlayAudioRequestPacket   = "PlayAudioRequest"
	TypePlayYoutubeRequestPacket = "PlayYoutubeRequest"
	TypeRunCommandRequestPacket  = "RunCommandRequest"
	TypeStreamingRequestPacker   = "StreamingRequest"
	TypeRequestActionTaskPacket  = "RequestActionTask"
	TypeTaskCancelRequest        = "TaskCancelRequest"
)

type RequestPacket struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}

type PlayVideoRequestPacket struct {
	URL     string `json:"url"`
	Volume  int    `json:"volume"`
	Visible bool   `json:"visible"`
}

type PlayYoutubeRequestPacket struct {
	URL     string `json:"url"`
	Volume  int    `json:"volume"`
	Visible bool   `json:"visible"`
}

type PlayAudioRequestPacket struct {
	URL    string `json:"url"`
	Volume int    `json:"volume"`
	Format string `json:"format"`
}

type RunCommandRequestPacket struct {
	Command string `json:"command"`
}

type StreamingRequestPacket struct {
	FrameRate int    `json:"frameRate"`
	Channel   string `json:"channel"`
}

type RequestTaskActionPacket struct {
	Task       string      `json:"task"`
	ActionType string      `json:"actionType"`
	Action     interface{} `json:"action"`
}

type TaskCancelRequest struct {
	TaskId string `json:"taskId"`
}
