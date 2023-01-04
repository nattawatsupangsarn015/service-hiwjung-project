package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Ingredient struct {
	IngredientID bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Amount       float64       `json:"amount" bson:"amount"`
}

type Ingredients struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	IsActive  bool          `json:"isActive" bson:"isActive"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}
