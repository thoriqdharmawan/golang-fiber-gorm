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
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "error query:"+err.Error())
	}

	var total int64
	if err := database.DB.Model(&entity.User{}).Count(&total).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "error get count"+err.Error())
	}

	meta := utils.GenerateMetaData(total, limit, offset)

	return utils.SuccessResponseWithMeta(ctx, fiber.StatusOK, users, meta)
}

func UserHandleCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
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
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Internal Server Error")
	}

	return utils.SuccessResponse(ctx, 200, newUser)
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User

	if err := database.DB.Preload("Posts").Preload("Language").Where("id = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(ctx, fiber.StatusNotFound, "User not found")
		}
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Internal Server Error")
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
		Language:  user.Language,
	}

	return utils.SuccessResponse(ctx, 200, userResponse)
}

func UserHandlerUpdateById(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	userId := ctx.Params("id")

	var user entity.User
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(ctx, fiber.StatusNotFound, "User not found")
		}
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Internal Server Error")
	}

	if err := database.DB.Model(&user).Updates(entity.User{
		Name:    userRequest.Name,
		Address: userRequest.Address,
		Phone:   userRequest.Phone,
	}).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, user)
}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserEmailRequest)

	if err := ctx.BodyParser(userRequest); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()

	if err := validate.Struct(userRequest); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	userId := ctx.Params("id")

	var user entity.User
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(ctx, fiber.StatusNotFound, "User not found")
		}
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Internal Server Error")
	}

	var userIsEmailExists entity.User
	if result := database.DB.Not("id = ?", userId).Where("email = ?", userRequest.Email).First(&userIsEmailExists); result.Error == nil {
		return utils.ErrorResponse(ctx, fiber.StatusForbidden, "Email already exists")
	}

	if err := database.DB.Model(&user).Update("email", userRequest.Email).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, user)
}

func UserHandlerDeleteUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User

	if err := database.DB.First(&user, userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(ctx, fiber.StatusNotFound, "User not found")
		}
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Internal Server Error")
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to delete user")
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, user)
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

func UserSetLanguage(ctx *fiber.Ctx) error {
	requestData := new(request.UserSetLanguageRequest)

	if err := ctx.BodyParser(requestData); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "error parser: "+err.Error())
	}

	validate := validator.New()

	if err := validate.Struct(requestData); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "error stuct validation: "+err.Error())
	}

	var user entity.User
	var language entity.Language

	if err := database.DB.First(&user, requestData.UserId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(ctx, fiber.StatusNotFound, "User not found: "+err.Error())
		}
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "internal server error: "+err.Error())
	}

	if err := database.DB.First(&language, requestData.LanguageId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorResponse(ctx, fiber.StatusNotFound, "Language not found: "+err.Error())
		}
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "internal server error: "+err.Error())
	}

	if err := database.DB.Model(&user).Association("Language").Append(&language); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "internal server error association: "+err.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, fiber.Map{
		"user":     user,
		"language": language,
	})
}
