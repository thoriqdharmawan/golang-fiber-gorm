package handler

import (
	"fmt"
	"golang-fiber-gorm/database"
	"golang-fiber-gorm/model/entity"
	"golang-fiber-gorm/model/request"
	"golang-fiber-gorm/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookHandlerGetAll(ctx *fiber.Ctx) error {
	var books []entity.Book

	limit := ctx.QueryInt("limit", 10)
	offset := ctx.QueryInt("offset", 0)

	if err := database.DB.Limit(limit).Offset(offset).Find(&books).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	var total int64
	if err := database.DB.Model(&entity.Book{}).Count(&total).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	meta := utils.GenerateMetaData(total, limit, offset)

	return utils.SuccessResponseWithMeta(ctx, fiber.StatusOK, books, meta)
}

func BookHandleCreate(ctx *fiber.Ctx) error {
	title := ctx.FormValue("title")
	author := ctx.FormValue("author")
	file, errorFile := ctx.FormFile("cover")

	book := request.BookCreateRequest{
		Title:  title,
		Author: author,
		Cover:  "",
	}

	validate := validator.New()

	if err := validate.Struct(book); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}

	if errorFile != nil {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Cover is required")
	}

	// 2mb
	if file.Size > 2*1024*1024 {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Cover size should not exceed 2MB")
	}

	// jpeg or png only
	if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
		return utils.ErrorResponse(ctx, fiber.StatusBadRequest, "Cover must be a JPEG or PNG image")
	}

	if err := ctx.SaveFile(file, fmt.Sprintf("./public/cover/%s", file.Filename)); err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	newBook := entity.Book{
		Title:  title,
		Author: author,
		Cover:  file.Filename,
	}

	if err := database.DB.Create(&newBook).Error; err != nil {
		return utils.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(ctx, fiber.StatusOK, newBook)
}
