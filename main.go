package main

import (
	"log"
	"xendit/db"
	"xendit/handlers"
	"xendit/repository"
	"xendit/routes"
	"xendit/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	client := db.ConnectDB()

	defer db.CloseDB()

	// repository
	repoOrg := repository.NewOrganizationRepository(client)
	repoComment := repository.NewCommentRepository(client)

	// service
	serviceOrg := services.NewOrganizationService(repoOrg)
	serviceComment := services.NewCommentService(repoComment, repoOrg)

	// handler
	handlerOrg := handlers.NewOrganizationHandler(serviceOrg)
	handlerComment := handlers.NewCommentHandler(serviceComment)

	// route
	routeOrg := routes.NewOrganizationRoute(handlerOrg)
	routeComment := routes.NewCommentRoute(handlerComment)

	routeOrg.OrgRoute(app)
	routeComment.RouteComment(app)

	log.Fatal(app.Listen(""))

}
