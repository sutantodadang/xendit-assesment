package handlers

import (
	"xendit/helpers"
	"xendit/models"
	"xendit/services"
	"xendit/utils"

	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	service *services.CommentService
}

func NewCommentHandler(service *services.CommentService) *CommentHandler {
	return &CommentHandler{service}
}

func (h *CommentHandler) CreateCommentHandler(c *fiber.Ctx) error {

	param := c.Params("organization")

	if param == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "parameter not found",
		})
	}

	comment := new(models.InputComment)

	if err := c.BodyParser(comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := utils.ValidateStruct(comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := h.service.CreateCommentService(*comment, param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "comment successfully added",
	})
}

func (h *CommentHandler) GetAllComment(c *fiber.Ctx) error {

	param := c.Params("organization")

	if param == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "parameter not found",
		})
	}

	res, err := h.service.GetAllCommentByOrg(param)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	response := helpers.ApiResponse("success", fiber.StatusOK, res)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *CommentHandler) DeleteAllHandler(c *fiber.Ctx) error {

	param := c.Params("organization")

	if param == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "parameter not found",
		})
	}

	if err := h.service.DeleteAllByOrg(param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "successfully deleted",
	})

}
