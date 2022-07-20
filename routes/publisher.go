package routes

import (
	"bookstore/handlers"
	"github.com/gofiber/fiber/v2"
)

func PublisherRoutes(app fiber.Router) {
	api := app.Group("/publisher")

	api.Post("/", handlers.CreatePublisher)
}
