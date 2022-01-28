package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organization struct {
	Id        primitive.ObjectID `json:"id" bson:"_id" `
	Name      string             `json:"name" bson:"name"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
