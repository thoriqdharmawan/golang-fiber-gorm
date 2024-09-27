package handler

import (
	"golang-fiber-gorm/database"
	"golang-fiber-gorm/model/entity"
	"golang-fiber-gorm/model/request"
	"golang-fiber-gorm/model/response"
	"golang-fiber-gorm/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	limit := ctx.QueryInt("limit", 10)
	offset := ctx.QueryInt("offset", 0)

	if err := database.DB.Preload("Posts").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return utils.ErrorResponse(ctx, 500, "error query:"+err.Error())
	}

	var total int64
	if err := database.DB.Model(&entity.User{}).Count(&total).Error; err != nil {
		return utils.ErrorResponse(ctx, 500, "error get count"+err.Error())
	}

	meta := utils.GenerateMetaData(total, limit, offset)

	return utils.SuccessResponseWithMeta(ctx, 200, users, meta)
}

func UserHandleCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return utils.ErrorResponse(ctx, 400, err.Error())
	}

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		return utils.ErrorResponse(ctx, 400, err.Error())
	}

	var userIsEmailExists entity.User
	if result := database.DB.Where("email = ?", user.Email).First(&userIsEmailExists); result.Error == nil {
		return utils.ErrorResponse(ctx, fiber.StatusForbidden, "Email already exists")
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Phone:    user.Phone,
		Password: hashedPassword,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return utils.ErrorResponse(ctx, 500, "Internal Server Error")
	}

	return utils.SuccessResponse(ctx, 200, newUser)

}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User

	if err := database.DB.Preload("Posts").Where("id = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(ctx, 404, "User not found")
		}
		return utils.ErrorResponse(ctx, 500, "Internal Server Error")
	}

	userResponse := response.UserResponseById{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Address:   user.Address,
		Phone:     user.Phone,
		Posts:     user.Posts,
	}

	return utils.SuccessResponse(ctx, 200, userResponse)
}

func UserHandlerUpdateById(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return utils.ErrorResponse(ctx, 400, err.Error())
	}

	userId := ctx.Params("id")

	var user entity.User
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(ctx, 404, "User not found")
		}
		return utils.ErrorResponse(ctx, 500, "Internal Server Error")
	}

	if err := database.DB.Model(&user).Updates(entity.User{
		Name:    userRequest.Name,
		Address: userRequest.Address,
		Phone:   userRequest.Phone,
	}).Error; err != nil {
		return utils.ErrorResponse(ctx, 500, err.Error())
	}

	return utils.SuccessResponse(ctx, 200, user)
}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserEmailRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return utils.ErrorResponse(ctx, 400, err.Error())
	}

	validate := validator.New()

	if err := validate.Struct(userRequest); err != nil {
		return utils.ErrorResponse(ctx, 400, err.Error())
	}

	userId := ctx.Params("id")

	var user entity.User
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(ctx, 404, "User not found")
		}
		return utils.ErrorResponse(ctx, 500, "Internal Server Error")
	}

	var userIsEmailExists entity.User
	if result := database.DB.Not("id = ?", userId).Where("email = ?", userRequest.Email).First(&userIsEmailExists); result.Error == nil {
		return utils.ErrorResponse(ctx, 403, "Email already exists")
	}

	if err := database.DB.Model(&user).Update("email", userRequest.Email).Error; err != nil {
		return utils.ErrorResponse(ctx, 500, err.Error())
	}

	return utils.SuccessResponse(ctx, 200, user)
}

func UserHandlerDeleteUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User

	if err := database.DB.First(&user, userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(ctx, 404, "User not found")
		}
		return utils.ErrorResponse(ctx, 500, "Internal Server Error")
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return utils.ErrorResponse(ctx, 500, "Failed to delete user")
	}

	return utils.SuccessResponse(ctx, 200, user)
}

func UserHandlerGetPosts(ctx *fiber.Ctx) error {
	var users []response.UserResponseGetPosts

	limit := ctx.QueryInt("limit", 10)
	offset := ctx.QueryInt("offset", 0)

	if err := database.DB.Model(&entity.User{}).
		Preload("Posts").
		Limit(limit).
		Offset(offset).
		Find(&users).
		Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "error query: "+err.Error())
	}

	var total int64
	if err := database.DB.Model(&entity.User{}).Count(&total).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "error get count"+err.Error())
	}

	meta := utils.GenerateMetaData(total, limit, offset)

	return utils.SuccessResponseWithMeta(ctx, 200, users, meta)
}
