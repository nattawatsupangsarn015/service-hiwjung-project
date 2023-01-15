package model

type ReplyMessage struct {
	ReplyToken string `json:"replyToken" bson:"replyToken"`
	Messages   []Text `json:"messages" bson:"messages"`
}

type Text struct {
	Type string `json:"type" bson:"type"`
	Text string `json:"text" bson:"text"`
}
