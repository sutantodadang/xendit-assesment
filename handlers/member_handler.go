package handlers

import (
	"xendit/helpers"
	"xendit/models"
	"xendit/services"
	"xendit/utils"

	"github.com/gofiber/fiber/v2"
)

type MemberHandler struct {
	service *services.MemberService
}

func NewMemberHandler(handler *services.MemberService) *MemberHandler {
	return &MemberHandler{handler}
}

func (h *MemberHandler) CreateMemberHandler(c *fiber.Ctx) error {

	param := c.Params("organization")

	if param == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "parameter not found",
		})
	}

	member := new(models.Member)

	if err := c.BodyParser(member); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := utils.ValidateStruct(member); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := h.service.CreateMember(*member, param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "successfully created",
	})
}

func (h *MemberHandler) GetAllMemberHandler(c *fiber.Ctx) error {

	param := c.Params("organization")

	if param == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "parameter not found",
		})
	}

	res, err := h.service.FindAllMemberByOrg(param)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	response := helpers.ApiResponse("success", fiber.StatusOK, res)

	return c.Status(fiber.StatusOK).JSON(response)
}
