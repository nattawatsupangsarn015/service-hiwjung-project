package responses

type LineMessage struct {
	Destination string        `json:"destination" bson:"destination"`
	Events      []EventStruct `json:"events" bson:"events"`
}

type EventStruct struct {
	ReplyToken string            `json:"replyToken"`
	Type       string            `json:"type"`
	Timestamp  int64             `json:"timestamp"`
	Source     SourceLineStruct  `json:"source"`
	Message    MessageLineStruct `json:"message"`
}

type SourceLineStruct struct {
	Type   string `json:"type"`
	UserID string `json:"userId"`
}

type MessageLineStruct struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}
