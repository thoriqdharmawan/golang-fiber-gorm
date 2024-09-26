package middleware

import (
	"golang-fiber-gorm/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")

	if token != "secret" {
		return utils.ErrorResponse(ctx, fiber.StatusUnauthorized, "Unauthorized")
	}

	return ctx.Next()
}
