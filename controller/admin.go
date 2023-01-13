package controller

import (
	"context"
	"example/service-hiwjung-project/config"
	"example/service-hiwjung-project/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var AdminCollection *mongo.Collection = config.GetCollection(config.DB, "admins")

func CreateAdmin(admin *model.Admin) error {
	err := CreateUniqueField(AdminCollection, bson.M{"username": 1})
	if err != nil {
		return err
	}

	_, err = AdminCollection.InsertOne(context.TODO(), admin)
	return err
}

func GetAdminByUsername(username string) (interface{}, error) {
	filter := bson.M{"username": username}
	var admin model.Admin
	err := AdminCollection.FindOne(context.TODO(), filter).Decode(&admin)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return admin, nil
}
