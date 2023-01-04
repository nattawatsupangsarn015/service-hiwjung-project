package controller

import (
	"example/service-hiwjung-project/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = config.GetCollection(config.DB, "menus")

func GetAllMenus() string {
	return "Get all menus"
}

func CreateMenu() string {
	return "CREATED !!"
}
