package controller

import (
	"example/service-hiwjung-project/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var ingredientCollection *mongo.Collection = config.GetCollection(config.DB, "ingredients")

func GetAllIngredients() string {
	return "Get all ingredient"
}

func CreateIngredient() string {
	return "CREATED Ingredient!!"
}
