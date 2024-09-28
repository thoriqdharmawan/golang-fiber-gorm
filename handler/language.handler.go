package handler

import (
	"golang-fiber-gorm/database"
	"golang-fiber-gorm/model/entity"
	"golang-fiber-gorm/model/request"
	"golang-fiber-gorm/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LanguageHandlerGetAll(ctx *fiber.Ctx) error {
	var languages []entity.Language

	limit := ctx.QueryInt("limit", 10)
	offset := ctx.QueryInt("offset", 0)

	if err := database.DB.Limit(limit).Offset(offset).Find(&languages).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "error query:"+err.Error())
	}

	var total int64
	if err := database.DB.Model(&entity.Language{}).Count(&total).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "error get count"+err.Error())
	}

	meta := utils.GenerateMetaData(total, limit, offset)

	return utils.SuccessResponseWithMeta(ctx, fiber.StatusOK, languages, meta)
}

func LanguageHandlerCreate(ctx *fiber.Ctx) error {
	language := new(request.LanguageCreateRequest)

	if err := ctx.BodyParser(language); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "error body parser: "+err.Error())
	}

	validate := validator.New()

	if err := validate.Struct(language); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "error stuct validation: "+err.Error())
	}

	newLanguage := entity.Language{
		Name: language.Name,
	}

	if err := database.DB.Create(&newLanguage).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, "error create: "+err.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, language)
}
