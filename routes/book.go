package routes

import (
	"bookstore/handlers"
	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app fiber.Router) {
	api := app.Group("/book")

	api.Post("/", handlers.CreateBook)
}
