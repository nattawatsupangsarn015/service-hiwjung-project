package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dishes struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Rating      float64            `json:"rating" bson:"rating"`
	Ingredients []Ingredient       `json:"ingredients" bson:"ingredients"`
	IsActive    bool               `json:"isActive" bson:"isActive"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
