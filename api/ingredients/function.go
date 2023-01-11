package ingredients

import (
	"example/service-hiwjung-project/controller"
	"example/service-hiwjung-project/model"
	"example/service-hiwjung-project/responses"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllIngredients() interface{} {
	return controller.GetAllIngredients()
}

func GetIngredientById(id string) interface{} {
	return controller.GetIngredientById(id)
}

func GetIngredientByName(name string) (interface{}, error) {
	result, err := controller.GetIngredientByName(name)
	if err != nil {
		return model.Ingredients{}, err
	}

	if result.IsActive != true {
		return bson.M{}, nil
	}

	fmt.Println(result)

	return result, nil
}

func CreateIngredient(Ingredient interface{}) (interface{}, error) {
	return Ingredient, nil
}

func UpdateIngredientById(id string, ingredient responses.IngredientRequestUpdate) (interface{}, error) {
	return controller.UpdateIngredientById(id, ingredient)
}
