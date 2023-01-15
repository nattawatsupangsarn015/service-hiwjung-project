package controller

import (
	"example/service-hiwjung-project/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = config.GetCollection(config.DB, "users")

// func CreateUser(user *model.Admin) error {
// 	err := CreateUniqueField(AdminCollection, bson.M{"username": 1})
// 	if err != nil {
// 		return err
// 	}

// 	_, err = AdminCollection.InsertOne(context.TODO(), admin)
// 	return err
// }
