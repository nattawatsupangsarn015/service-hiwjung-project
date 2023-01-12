package ingredients

import (
	"encoding/json"
	"errors"
	"example/service-hiwjung-project/controller"
	"example/service-hiwjung-project/model"
	"example/service-hiwjung-project/responses"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllIngredients() interface{} {
	return controller.GetAllIngredients()
}

func GetIngredientById(id string) interface{} {
	return controller.GetIngredientById(id)
}

func GetIngredientByName(name string) (interface{}, error) {
	var ingredient model.Ingredients
	result, err := controller.GetIngredientByName(name)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return bson.M{}, nil
	} else {
		if convertModel, ok := result.(model.Ingredients); ok {
			if convertModel.IsActive == true {
				return convertModel, nil
			} else {
				return bson.M{}, nil
			}
		} else {
			jsonBytes, ok := result.([]byte)
			if !ok {
				return nil, errors.New("Cannot convert data to byte.")
			}

			err = json.Unmarshal(jsonBytes, &ingredient)
			if err != nil {
				return nil, err
			}

			if ingredient.IsActive != true {
				return bson.M{}, nil
			}

			return result, nil
		}
	}
}

func CreateIngredient(Ingredient interface{}) (interface{}, error) {
	return Ingredient, nil
}

func CreateIngredientActive(name string) (interface{}, error) {
	var findIngredient model.Ingredients
	result, err := controller.GetIngredientByName(name)

	if result != nil {
		if convertResult, ok := result.(model.Ingredients); ok {
			if convertResult.IsActive == false {
				err = ActivateIngredientById(convertResult.ID.Hex())
				if err != nil {
					return convertResult.ID, nil
				}
				return convertResult.ID, nil
			} else {
				return convertResult.ID, nil
			}
		} else {
			jsonBytes, ok := result.([]byte)
			if !ok {
				return nil, errors.New("Cannot convert data to byte.")
			}

			err = json.Unmarshal(jsonBytes, &findIngredient)
			if err != nil {
				return nil, err
			}

			if findIngredient.IsActive == false {
				err = ActivateIngredientById(findIngredient.ID.Hex())
				if err != nil {
					return nil, err
				}

				return findIngredient.ID, nil
			} else {
				return findIngredient.ID, nil
			}
		}
	} else {
		newIngredient := model.Ingredients{
			ID:        primitive.NewObjectID(),
			Name:      name,
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		resultId, err := controller.CreateIngredientActive(newIngredient)
		if err != nil {
			return nil, err
		}

		return resultId, nil
	}
}

func UpdateIngredientById(id string, ingredient responses.IngredientRequestUpdate) (interface{}, error) {
	return controller.UpdateIngredientById(id, ingredient)
}

func ActivateIngredientById(id string) error {
	err := controller.ActivateIngredientById(id)
	if err != nil {
		return err
	}
	return nil
}
