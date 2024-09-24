package handler

import (
	"golang-fiber-gorm/database"
	helpers "golang-fiber-gorm/helper"
	"golang-fiber-gorm/model/entity"
	"golang-fiber-gorm/model/request"
	"golang-fiber-gorm/model/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	if err := database.DB.Find(&users).Error; err != nil {
		return helpers.ErrorResponse(ctx, 500, "Internal Server Error")
	}

	return helpers.SuccessResponse(ctx, 200, users)
}

func UserHandleCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return helpers.ErrorResponse(ctx, 400, err.Error())
	}

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		return helpers.ErrorResponse(ctx, 400, err.Error())
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return helpers.ErrorResponse(ctx, 500, "Internal Server Error")
	}

	return helpers.SuccessResponse(ctx, 200, newUser)

}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User

	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return helpers.ErrorResponse(ctx, 404, "Not found")
	}

	userResponse := response.UserResponseById{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return helpers.SuccessResponse(ctx, 200, userResponse)
}
