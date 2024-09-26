package handler

import (
	"golang-fiber-gorm/database"
	"golang-fiber-gorm/model/entity"
	"golang-fiber-gorm/model/request"
	"golang-fiber-gorm/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(loginRequest); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	var user entity.User
	if err := database.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(ctx, fiber.StatusForbidden, "Wrong Email or Password")
		}
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Internal Server Error")
	}

	if isPasswordMatch := utils.CheckPasswordHash(loginRequest.Password, user.Password); !isPasswordMatch {
		return utils.ErrorResponse(ctx, fiber.StatusForbidden, "Wrong Email or Password")
	}

	token, errGenerateToken := utils.GenerateJWTToken(user)

	if errGenerateToken != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, errGenerateToken.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, fiber.Map{
		"token": token,
	})
}
