package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Menu struct {
	ID         bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Rating     float64       `json:"rating" bson:"rating"`
	Ingredient []Ingredient  `json:"Ingredient" bson:"Ingredient"`
	IsActive   bool          `json:"isActive" bson:"isActive"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" bson:"updated_at"`
}
