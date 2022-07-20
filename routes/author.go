package routes

import (
	"bookstore/handlers"
	"github.com/gofiber/fiber/v2"
)

func AuthorRoutes(app fiber.Router) {
	api := app.Group("/author")

	api.Post("/", handlers.CreateAuthor)
}
