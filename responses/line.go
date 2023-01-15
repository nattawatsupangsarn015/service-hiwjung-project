package responses

type LineMessage struct {
	Destination string        `json:"destination" bson:"destination"`
	Events      []EventStruct `json:"events" bson:"events"`
}

type EventStruct struct {
	ReplyToken string            `json:"replyToken" bson:"replyToken"`
	Type       string            `json:"type" bson:"type"`
	Timestamp  int64             `json:"timestamp" bson:"timestamp"`
	Source     SourceLineStruct  `json:"source" bson:"source"`
	Message    MessageLineStruct `json:"message" bson:"message"`
}

type SourceLineStruct struct {
	Type   string `json:"type" bson:"type"`
	UserID string `json:"userId" bson:"userId"`
}

type MessageLineStruct struct {
	ID   string `json:"id" bson:"id"`
	Type string `json:"type" bson:"type"`
	Text string `json:"text" bson:"text"`
}
