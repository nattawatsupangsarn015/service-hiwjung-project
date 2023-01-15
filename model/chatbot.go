package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatbotReply struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Message   string             `json:"message" bson:"message"`
	Reply     string             `json:"reply" bson:"reply"`
	Type      string             `json:"type" bson:"type"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
