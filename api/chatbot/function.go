package chatbot

import (
	"example/service-hiwjung-project/controller"
	"example/service-hiwjung-project/model"
	"example/service-hiwjung-project/responses"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateChatbot(chatbot responses.RequestCreateChatbot) (interface{}, error) {
	findChatbot, err := controller.FindChatbotMessages(chatbot.Message)
	if err != nil {
		return nil, nil
	}

	if findChatbot != nil {
		return findChatbot, nil
	} else {
		newChatbot := model.ChatbotReply{
			ID:        primitive.NewObjectID(),
			Message:   chatbot.Message,
			Reply:     chatbot.Reply,
			Type:      chatbot.Type,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = controller.CreateChatbot(newChatbot)
		if err != nil {
			return nil, err
		}
		return newChatbot, nil
	}
}
