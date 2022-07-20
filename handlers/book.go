package handlers

import (
	"bookstore/database"
	"bookstore/models"
	"bookstore/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func CreateBook(ctx *fiber.Ctx) error {
	var err error
	book := new(models.Book)

	if err = ctx.BodyParser(book); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Corrupted Body.")
	}

	book.Name = strings.TrimSpace(book.Name)
	if book.Name == "" || book.Price == nil || book.AuthorID == nil || book.PublisherID == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Check Your Inputs.")
	}

	if err = database.DB.Create(book).Error; err != nil {
		return err
	}

	return utils.SuccessPresenter(ctx, "Book Created", book)
}
