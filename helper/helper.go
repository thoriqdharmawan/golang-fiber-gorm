package helpers

import "github.com/gofiber/fiber/v2"

func ErrorResponse(ctx *fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}

func SuccessResponse(ctx *fiber.Ctx, statusCode int, data interface{}) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}

