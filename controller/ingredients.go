package controller

import (
	"context"
	"example/service-hiwjung-project/config"
	"example/service-hiwjung-project/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ingredientCollection *mongo.Collection = config.GetCollection(config.DB, "ingredients")

func GetAllIngredients() interface{} {
	filter := bson.D{};
	cursor, err := ingredientCollection.Find(context.TODO(), filter);
	if err != nil {
		panic(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results;
}

func GetIngredientById(id string) interface{} {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objectId}}
	var result model.Ingredients;
    err := ingredientCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil;
		}
		panic(err)
	}
	return result;
}

func CreateIngredient() string {
	return "CREATED Ingredient!!"
}
