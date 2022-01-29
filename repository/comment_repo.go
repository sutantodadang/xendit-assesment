package repository

import (
	"context"
	"xendit/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICommentRepository interface {
	Save(comment models.Comment) error
	FindAll(org models.Organization) ([]models.Comment, error)
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

func (r *CommentRepository) FindAll(org models.Organization) ([]models.Comment, error) {

	res, err := r.db.Collection("comments").Find(context.Background(), bson.M{"organization": org})

	if err != nil {
		return nil, err
	}

	var result []models.Comment

	if err := res.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil

}
