package packets

type TaskHandShake struct {
	TaskType string `json:"taskType"`
}

type TaskHandshakeResponse struct {
	Id     string `json:"id"`
	TaskId string `json:"taskId"`
}
