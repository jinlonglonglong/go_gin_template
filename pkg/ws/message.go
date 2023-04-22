package ws

type Message struct {
	Type int32       `json:"type"`
	Data interface{} `json:"data"`
}
