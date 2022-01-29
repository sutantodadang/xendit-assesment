package repository

import (
	"context"
	"strings"
	"xendit/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IOrganizationRepository interface {
	Save(org models.Organization) error
	FindById(id primitive.ObjectID) (models.Organization, error)
	FindByName(name string) (models.Organization, error)
}

type OrganizationRepo struct {
	db *mongo.Database
}

func NewOrganizationRepository(db *mongo.Database) *OrganizationRepo {
	return &OrganizationRepo{db}
}

func (r *OrganizationRepo) Save(org models.Organization) error {

	_, err := r.db.Collection("organization").InsertOne(context.Background(), &org)

	if err != nil {
		return err
	}

	return nil
}

func (r *OrganizationRepo) FindById(id primitive.ObjectID) (models.Organization, error) {

	res := r.db.Collection("organization").FindOne(context.Background(), bson.M{"_id": id})

	var result models.Organization

	if err := res.Decode(&result); err != nil {
		return models.Organization{}, err
	}

	return result, nil

}

func (r *OrganizationRepo) FindByName(name string) (models.Organization, error) {

	res := r.db.Collection("organization").FindOne(context.Background(), bson.M{"name": strings.ToLower(name)})

	var result models.Organization

	if err := res.Decode(&result); err != nil {
		return models.Organization{}, err
	}

	return result, nil
}
