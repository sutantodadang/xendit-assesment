package services

import (
	"reflect"
	"time"
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

	if reflect.ValueOf(res).IsValid() {
		return fiber.NewError(fiber.StatusBadRequest, "Organization Data Already Created")
	}

	if err != nil {
		return err
	}

	org.CreatedAt = time.Now()
	org.UpdatedAt = time.Now()

	if err := s.repo.Save(org); err != nil {
		return err
	}

	return nil
}
