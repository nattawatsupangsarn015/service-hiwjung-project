package ingredients

import (
	"example/service-hiwjung-project/controller"
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
