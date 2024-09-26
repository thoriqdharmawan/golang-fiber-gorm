package middleware

import (
	helpers "golang-fiber-gorm/helper"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")

	if token != "secret" {
		return helpers.ErrorResponse(ctx, fiber.StatusUnauthorized, "Unauthorized")
	}

	return ctx.Next()
}
