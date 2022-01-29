package handlers

import (
	"xendit/models"
	"xendit/services"
	"xendit/utils"

	"github.com/gofiber/fiber/v2"
)

type OrganizationHandler struct {
	service *services.OrganizationService
}

func NewOrganizationHandler(service *services.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{service}
}

func (h *OrganizationHandler) CreateOrganizationHandler(c *fiber.Ctx) error {

	name := new(models.InputOrganization)

	if err := c.BodyParser(name); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := utils.ValidateStruct(name); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := h.service.CreateOrg(*name); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "organization successfull created",
	})
}
