package controller

import (
	"context"
	"example/service-hiwjung-project/config"
	"example/service-hiwjung-project/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DishCollection *mongo.Collection = config.GetCollection(config.DB, "dishes")

func CreateDish(dish *model.Dishes) error {
	err := CreateUniqueField(DishCollection)
	if err != nil {
		return err
	}

	_, err = DishCollection.InsertOne(context.TODO(), dish)
	return err
}

func CreateUniqueField(collection *mongo.Collection) error {
	_, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bson.M{"name": 1},
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		return err
	}

	return nil
}
