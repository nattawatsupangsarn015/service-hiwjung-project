package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Product struct {
	ProductID    bson.ObjectId `json:"product_id" bson:"_id,omitempty"`
	ProductName  string        `json:"product_name" bson:"product_name"`
	ProductPrice string        `json:"product_price" bson:"product_price"`
	Amount       int           `json:"amount" bson:"amount"`
	CreatedAt    time.Time     `json:"-" bson:"created_at"`
	UpdatedAt    time.Time     `json:"updated_time" bson:"updated_at"`
}
