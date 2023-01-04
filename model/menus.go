package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type RawMaterial struct {
	RawMaterialID bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Amount        float64       `json:"amount" bson:"amount"`
}

type Menu struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Rating      float64       `json:"rating" bson:"rating"`
	RawMaterial []RawMaterial `json:"rawMaterial" bson:"rawMaterial"`
	IsActive    bool          `json:"isActive" bson:"isActive"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}

type RawMaterials struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	IsActive  bool          `json:"isActive" bson:"isActive"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}
