package handler

import (
	"golang-fiber-gorm/database"
	"golang-fiber-gorm/model/entity"
	"golang-fiber-gorm/model/request"
	"golang-fiber-gorm/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PostHandlerCreate(ctx *fiber.Ctx) error {
	post := new(request.PostCreateRequest)

	if err := ctx.BodyParser(post); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "error body parser: "+err.Error())
	}

	validate := validator.New()

	if err := validate.Struct(post); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "error stuct validation: "+err.Error())
	}

	newPost := entity.Post{
		Title:  post.Title,
		UserID: post.UserID,
	}

	if err := database.DB.Create(&newPost).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "error create: "+err.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, post)
}

func PostHandlerGetAllPost(ctx *fiber.Ctx) error {
	var posts []entity.Post

	limit := ctx.QueryInt("limit", 10)
	offset := ctx.QueryInt("offset", 0)

	if err := database.DB.Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	var total int64
	if err := database.DB.Model(&entity.Post{}).Count(&total).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	meta := utils.GenerateMetaData(total, limit, offset)

	return utils.SuccessResponseWithMeta(ctx, 200, posts, meta)
}
