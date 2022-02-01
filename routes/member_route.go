package routes

import (
	"xendit/handlers"

	"github.com/gofiber/fiber/v2"
)

type MemberRoute struct {
	handler *handlers.MemberHandler
}

func NewMemberRoute(handler *handlers.MemberHandler) *MemberRoute {
	return &MemberRoute{handler}
}

func (rt *MemberRoute) RouteMember(app *fiber.App) {

	route := app.Group("/api/v1/orgs/:organization/member")

	route.Post("/", rt.handler.CreateMemberHandler)
	route.Get("/", rt.handler.GetAllMemberHandler)
}
