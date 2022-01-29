package services

import (
	"strings"
	"time"
	"xendit/models"
	"xendit/repository"

	"github.com/gofiber/fiber/v2"
)

type IOrganizationService interface {
	CreateOrg(name models.InputOrganization) error
}

type OrganizationService struct {
	repo *repository.OrganizationRepo
}

func NewOrganizationService(repo *repository.OrganizationRepo) *OrganizationService {
	return &OrganizationService{repo}
}

func (s *OrganizationService) CreateOrg(input models.InputOrganization) error {

	res, _ := s.repo.FindByName(input.Name)

	if strings.ToLower(res.Name) == strings.ToLower(input.Name) {
		return fiber.NewError(fiber.StatusBadRequest, "Organization Data Already Created")
	}

	var org models.Organization

	org.CreatedAt = time.Now()
	org.UpdatedAt = time.Now()
	org.Name = input.Name

	if err := s.repo.Save(org); err != nil {
		return err
	}

	return nil
}
