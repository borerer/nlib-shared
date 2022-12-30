package nlibshared

type WebSocketMessage struct {
	MessageID     string      `json:"message_id"`
	PairMessageID string      `json:"pair_message_id"`
	Type          string      `json:"type"`
	SubType       string      `json:"sub_type"`
	Timestamp     int64       `json:"timestamp"`
	Payload       interface{} `json:"payload"`
}
