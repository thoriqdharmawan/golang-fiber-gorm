package handler

import (
	"golang-fiber-gorm/database"
	helpers "golang-fiber-gorm/helper"
	"golang-fiber-gorm/model/entity"
	"golang-fiber-gorm/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(loginRequest); err != nil {
		return helpers.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	var user entity.User
	if err := database.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.ErrorResponse(ctx, fiber.StatusForbidden, "Wrong Email or Password")
		}
		return helpers.ErrorResponse(ctx, fiber.StatusInternalServerError, "Internal Server Error")
	}

	if isPasswordMatch := helpers.CheckPasswordHash(loginRequest.Password, user.Password); !isPasswordMatch {
		return helpers.ErrorResponse(ctx, fiber.StatusForbidden, "Wrong Email or Password")
	}

	return helpers.SuccessResponse(ctx, fiber.StatusOK, fiber.Map{
		"token": "secret-token",
	})
}
