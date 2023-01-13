package controller

import (
	"context"
	"example/service-hiwjung-project/config"
	"example/service-hiwjung-project/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var DishCollection *mongo.Collection = config.GetCollection(config.DB, "dishes")

func CreateDish(dish *model.Dishes) error {
	err := CreateUniqueField(DishCollection, bson.M{"name": 1})
	if err != nil {
		return err
	}

	_, err = DishCollection.InsertOne(context.TODO(), dish)
	return err
}
