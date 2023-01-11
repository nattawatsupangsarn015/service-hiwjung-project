package dishes

import (
	"example/service-hiwjung-project/api/ingredients"
	"example/service-hiwjung-project/model"
	"example/service-hiwjung-project/responses"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateDish(dish responses.DishRequestCreate) (interface{}, error) {
	newDish := &model.Dishes{
		ID:          primitive.NewObjectID(),
		Name:        dish.Name,
		Ingredients: dish.Ingredients,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	var wg sync.WaitGroup
	var ingredientsData []interface{}

	for i := range newDish.Ingredients {
		wg.Add(i)
		go func(i int) {
			defer wg.Done()
			ingredient, err := ingredients.GetIngredientByName(newDish.Ingredients[i])
			if err != nil {
				return
			}
			ingredientsData = append(ingredientsData, ingredient)
		}(i)
	}

	wg.Wait()
	fmt.Println(ingredientsData)

	// err := controller.CreateDish(newDish)
	// if err != nil {
	// 	return nil, err
	// }
	// return newDish, nil

	return newDish, nil
}
