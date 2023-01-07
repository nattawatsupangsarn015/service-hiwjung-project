package ingredients

import (
	"example/service-hiwjung-project/controller"
	"example/service-hiwjung-project/responses"
)

func GetAllIngredients() interface{} {
	return controller.GetAllIngredients()
}

func GetIngredientById(id string) interface{} {
	return controller.GetIngredientById(id)
}

func CreateIngredient(Ingredient interface{}) (interface{}, error) {
	return Ingredient, nil
}

func UpdateIngredientById(id string, ingredient responses.IngredientRequestUpdate) (interface{}, error) {
	return controller.UpdateIngredientById(id, ingredient)
}
