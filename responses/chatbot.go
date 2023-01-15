package responses

type RequestCreateChatbot struct {
	Message string `json:"message" bson:"message"`
	Reply   string `json:"reply" bson:"reply"`
	Type    string `json:"type" bson:"type"`
}
