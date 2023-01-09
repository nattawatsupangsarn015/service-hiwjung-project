package responses

import "go.mongodb.org/mongo-driver/bson/primitive"

type IngredientRequestCreate struct {
	Name string `json:"name"`
}

type IngredientRequestUpdate struct {
	Name string `json:"name"`
}

type IngredientQueryById struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
}
