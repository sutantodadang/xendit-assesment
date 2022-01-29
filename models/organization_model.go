package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organization struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty" `
	Name      string             `json:"name" bson:"name" form:"name"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type InputOrganization struct {
	Name string `json:"name" validate:"required,min=3" bson:"name"`
}
