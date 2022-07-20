package handlers

import (
	"bookstore/database"
	"bookstore/models"
	"bookstore/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func CreatePublisher(ctx *fiber.Ctx) error {
	var err error
	publisher := new(models.Publisher)

	if err = ctx.BodyParser(publisher); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Corrupted Body.")
	}

	publisher.Name = strings.TrimSpace(publisher.Name)
	if publisher.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Check Your Inputs.")
	}

	if err = database.DB.Create(publisher).Error; err != nil {
		return err
	}

	return utils.SuccessPresenter(ctx, "Publisher Created.", publisher)
}
