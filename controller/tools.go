package controller

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUniqueField(collection *mongo.Collection, key interface{}) error {
	_, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    key,
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		return err
	}

	return nil
}
