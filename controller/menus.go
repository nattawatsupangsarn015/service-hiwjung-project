package controller

import (
	"context"
	"example/service-hiwjung-project/config"
	"example/service-hiwjung-project/responses"

	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = config.GetCollection(config.DB, "menus")

func GetAllMenus() string {
	return "Get all menus"
}

func CreateMenu(menu responses.MenuRequestCreate) error {
	_, err := menuCollection.InsertOne(context.TODO(), menu)
	return err
}
