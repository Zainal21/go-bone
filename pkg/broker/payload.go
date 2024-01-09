package broker

type MessagePayload struct {
	Type    string `json:"type"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}
