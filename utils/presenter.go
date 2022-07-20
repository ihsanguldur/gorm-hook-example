package utils

import "github.com/gofiber/fiber/v2"

func SuccessPresenter(ctx *fiber.Ctx, message string, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func ErrorPresenter(ctx *fiber.Ctx, message string, code int) error {
	return ctx.Status(code).JSON(fiber.Map{
		"success": false,
		"message": message,
		"data":    nil,
	})
}
