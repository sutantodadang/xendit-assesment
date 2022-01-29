package routes

import (
	"xendit/handlers"

	"github.com/gofiber/fiber/v2"
)

type OrganizationRoute struct {
	handler *handlers.OrganizationHandler
}

func NewOrganizationRoute(handler *handlers.OrganizationHandler) *OrganizationRoute {
	return &OrganizationRoute{handler}
}

func (rt *OrganizationRoute) OrgRoute(app *fiber.App) {

	route := app.Group("/api/v1/orgs")

	route.Post("/", rt.handler.CreateOrganizationHandler)
}
