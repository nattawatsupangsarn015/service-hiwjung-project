package controller

import (
	"context"
	"example/service-hiwjung-project/config"
	"example/service-hiwjung-project/model"
	"example/service-hiwjung-project/responses"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ingredientCollection *mongo.Collection = config.GetCollection(config.DB, "ingredients")

func GetAllIngredients() interface{} {
	filter := bson.D{}
	cursor, err := ingredientCollection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}

func GetIngredientById(id string) interface{} {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	var result model.Ingredients
	err := ingredientCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil
		}
		panic(err)
	}
	return result
}

func CreateIngredient(Ingredient responses.IngredientRequestCreate) error {
	_, err := ingredientCollection.InsertOne(context.TODO(), Ingredient)
	return err
}

func UpdateIngredientById(id string, ingredinet responses.IngredientRequestUpdate) (model.Ingredients, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return model.Ingredients{}, err
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": ingredinet}

	var updatedIngredient model.Ingredients
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = ingredientCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedIngredient)

	if err != nil {
		return model.Ingredients{}, err
	}

	return updatedIngredient, nil
}
