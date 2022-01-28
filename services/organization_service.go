package services

import (
	"reflect"
	"xendit/models"
	"xendit/repository"

	"github.com/gofiber/fiber/v2"
)

type IOrganizationService interface {
	CreateOrg(org models.Organization) error
}

type OrganizationService struct {
	repo *repository.OrganizationRepo
}

func NewOrganizationService(repo *repository.OrganizationRepo) *OrganizationService {
	return &OrganizationService{repo}
}

func (s *OrganizationService) CreateOrg(org models.Organization) error {

	res, err := s.repo.Find(org.Id)

	if reflect.ValueOf(res).IsNil() {
		return fiber.NewError(fiber.StatusNotFound, "Not Found Organization Data")
	}

	if err != nil {
		return err
	}

	if err := s.repo.Save(org); err != nil {
		return err
	}

	return nil
}
