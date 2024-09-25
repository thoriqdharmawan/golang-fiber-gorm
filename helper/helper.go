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

func SuccessResponseWithMeta(ctx *fiber.Ctx, statusCode int, data interface{}, meta interface{}) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"status": "success",
		"data":   data,
		"meta":   meta,
	})
}

func GenerateMetaData(total int64, limit int, offset int) fiber.Map {
	hasNextPage := (offset + limit) < int(total)
	hasPrevPage := offset > 0

	return fiber.Map{
		"total":       total,
		"limit":       limit,
		"offset":      offset,
		"hasNextPage": hasNextPage,
		"hasPrevPage": hasPrevPage,
	}
}
