package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Message      string             `json:"message" bson:"message"`
	Organization Organization       `json:"organization" bson:"organization"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

type InputComment struct {
	Message string `json:"message" bson:"message"`
}
