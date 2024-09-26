package utils

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
