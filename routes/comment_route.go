package routes

import (
	"xendit/handlers"

	"github.com/gofiber/fiber/v2"
)

type CommentRoute struct {
	handler *handlers.CommentHandler
}

func NewCommentRoute(handler *handlers.CommentHandler) *CommentRoute {
	return &CommentRoute{handler}
}

func (rt *CommentRoute) RouteComment(app *fiber.App) {

	route := app.Group("/api/v1/orgs/:organization/comment")

	route.Post("/", rt.handler.CreateCommentHandler)
	route.Get("/", rt.handler.GetAllComment)
}
