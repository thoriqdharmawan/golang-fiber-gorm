package middleware

import (
	"golang-fiber-gorm/utils"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	
	if err := utils.VerifyJWTTokenHandler(token); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusUnauthorized, err.Error())
	}

	return ctx.Next()
}
