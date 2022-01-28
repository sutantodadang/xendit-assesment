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

	repo := repository.NewOrganizationRepository(client)
	service := services.NewOrganizationService(repo)
	handler := handlers.NewOrganizationHandler(service)

	org := routes.NewOrganizationRoute(handler)

	org.OrgRoute(app)

	log.Fatal(app.Listen(""))

}
