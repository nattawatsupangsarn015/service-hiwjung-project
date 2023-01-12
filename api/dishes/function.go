package dishes

import (
	"example/service-hiwjung-project/api/ingredients"
	"example/service-hiwjung-project/controller"
	"example/service-hiwjung-project/model"
	"example/service-hiwjung-project/responses"
	"fmt"
	"reflect"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDish(dish responses.DishRequestCreate) (interface{}, error) {
	var ingredientModel []model.Ingredient
	var wg sync.WaitGroup
	for i := range dish.Ingredients {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defaultInterface := bson.M{}
			ingredient, err := ingredients.GetIngredientByName(dish.Ingredients[i].Value)
			if err != nil {
				panic(err)
			}

			if reflect.DeepEqual(ingredient, defaultInterface) || ingredient == nil {
				resultId, err := ingredients.CreateIngredientActive(dish.Ingredients[i].Value)
				if err != nil {
					panic(err)
				}

				str := fmt.Sprintf("%v", resultId)
				objectId, err := primitive.ObjectIDFromHex(str)
				if err != nil {
					panic(err)
				}

				ingredientData := model.Ingredient{
					IngredientID: objectId,
					Unit:         dish.Ingredients[i].Unit,
					Amount:       dish.Ingredients[i].Amount,
				}
				ingredientModel = append(ingredientModel, ingredientData)
			} else {
				convertIngredient, _ := ingredient.(model.Ingredients)
				ingredientData := model.Ingredient{
					IngredientID: convertIngredient.ID,
					Unit:         dish.Ingredients[i].Unit,
					Amount:       dish.Ingredients[i].Amount,
				}
				ingredientModel = append(ingredientModel, ingredientData)
			}
		}(i)
	}

	wg.Wait()

	newDish := &model.Dishes{
		ID:          primitive.NewObjectID(),
		Name:        dish.Name,
		Ingredients: ingredientModel,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := controller.CreateDish(newDish)
	if err != nil {
		return nil, err
	}
	return newDish, nil
}
