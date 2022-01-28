package repository

import (
	"context"
	"xendit/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ICommentRepository interface {
	Save(comment models.Comment) error
}

type CommentRepository struct {
	db *mongo.Database
}

func NewCommentRepository(db *mongo.Database) *CommentRepository {
	return &CommentRepository{db}
}

func (r *CommentRepository) Save(comment models.Comment) error {

	_, err := r.db.Collection("comments").InsertOne(context.Background(), &comment)

	if err != nil {
		return err
	}

	return nil
}
