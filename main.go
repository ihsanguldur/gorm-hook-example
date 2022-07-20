package main

import (
	"bookstore/database"
	"bookstore/routes"
	"bookstore/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	database.Connect()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
