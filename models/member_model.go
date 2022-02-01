package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Followers    int64              `json:"followers" bson:"followers"`
	Following    int64              `json:"following" bson:"following"`
	Avatar       string             `json:"avatar" bson:"avatar"`
	Name         string             `json:"name" validate:"required,min=3" bson:"name"`
	Organization Organization       `json:"organization" bson:"organization"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}
