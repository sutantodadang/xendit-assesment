package repository

import (
	"context"
	"xendit/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMemberRepository interface {
	Save(member models.Member) error
	FindAllByOrg(org models.Organization) ([]models.Member, error)
	FindByName(name string) (models.Member, error)
}

type MemberRepository struct {
	db *mongo.Database
}

func NewMemberRepository(db *mongo.Database) *MemberRepository {
	return &MemberRepository{db}
}

func (r *MemberRepository) Save(member models.Member) error {

	_, err := r.db.Collection("members").InsertOne(context.Background(), &member)

	if err != nil {
		return err
	}

	return nil
}

func (r *MemberRepository) FindAllByOrg(org models.Organization) ([]models.Member, error) {

	res, err := r.db.Collection("members").Find(context.Background(), bson.M{"organization": org})

	if err != nil {
		return nil, err
	}

	var result []models.Member

	if err := res.All(context.Background(), &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fiber.ErrNotImplemented
	}

	return result, nil

}

func (r *MemberRepository) FindByName(name string) (models.Member, error) {

	res := r.db.Collection("members").FindOne(context.Background(), bson.M{"name": name})

	var result models.Member

	if err := res.Decode(&result); err != nil {
		return models.Member{}, err
	}

	return result, nil

}
