package controller

import (
	"context"
	"example/service-hiwjung-project/config"
	"example/service-hiwjung-project/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var messagesCollection *mongo.Collection = config.GetCollection(config.DB, "messages")

func CreateChatbot(chatbot model.ChatbotReply) error {
	err := CreateUniqueField(messagesCollection, bson.M{"message": 1})
	if err != nil {
		return err
	}

	_, err = messagesCollection.InsertOne(context.TODO(), chatbot)
	return err
}

func FindChatbotMessages(message string) (interface{}, error) {
	filter := bson.M{"message": message}
	var result model.ChatbotReply
	err := messagesCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}
