package repository

import (
	"context"
	"xendit/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICommentRepository interface {
	Save(comment models.Comment) error
	FindAll(org models.Organization) ([]models.Comment, error)
	DeleteAll(org models.Organization) error
	SaveDump(comments []models.Comment) error
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

	if len(result) == 0 {
		return nil, fiber.ErrNotFound
	}

	return result, nil

}

func (r *CommentRepository) DeleteAll(org models.Organization) error {

	_, err := r.db.Collection("comments").DeleteMany(context.Background(), bson.M{"organization": org})

	if err != nil {
		return err
	}

	return nil
}

func (r *CommentRepository) SaveDump(comments []interface{}) error {

	_, err := r.db.Collection("dumps").InsertMany(context.Background(), comments)

	if err != nil {
		return err
	}

	return nil
}
