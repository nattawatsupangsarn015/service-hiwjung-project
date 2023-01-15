package controller

import (
	"context"
	"example/service-hiwjung-project/config"
	"example/service-hiwjung-project/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")

func CreateUser(user *model.Users) error {
	err := CreateUniqueField(userCollection, bson.M{"lineId": 1})
	if err != nil {
		return err
	}

	_, err = userCollection.InsertOne(context.TODO(), user)
	return err
}

func FindUserByLineId(lineId string) (interface{}, error) {
	filter := bson.M{"lineId": lineId}
	var result model.Users
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}
