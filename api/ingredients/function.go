package ingredients

import (
	"example/service-hiwjung-project/controller"
)

func GetAllIngredients() string {
	return controller.GetAllIngredients()
}

func CreateIngredient(Ingredient interface{}) (interface{}, error) {
	return Ingredient, nil
}
