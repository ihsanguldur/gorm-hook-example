package utils

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Something Gone Wrong."

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	} else if strings.Contains(err.Error(), "1452") {
		code = 400
		message = "Author or Publisher Does Not Found."
	}
	return ErrorPresenter(ctx, message, code)
}
