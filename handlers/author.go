package handlers

import (
	"bookstore/database"
	"bookstore/models"
	"bookstore/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func CreateAuthor(ctx *fiber.Ctx) error {
	var err error
	author := new(models.Author)

	if err = ctx.BodyParser(author); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Corrupted Body.")
	}

	author.FirstName = strings.TrimSpace(author.FirstName)
	author.LastName = strings.TrimSpace(author.LastName)
	if author.FirstName == "" || author.LastName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Check Your Inputs.")
	}

	if err = database.DB.Create(author).Error; err != nil {
		return err
	}

	return utils.SuccessPresenter(ctx, "Author Created.", author)
}
