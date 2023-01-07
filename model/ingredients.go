package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ingredient struct {
	IngredientID primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Amount       float64            `json:"amount" bson:"amount"`
}

type Ingredients struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	IsActive  bool               `json:"isActive" bson:"isActive"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
