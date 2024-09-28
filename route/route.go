package route

import (
	"golang-fiber-gorm/config"
	"golang-fiber-gorm/handler"
	"golang-fiber-gorm/middleware"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Static("/public", config.ProjectRootPath+"/public/asset/image1.avif")

	app.Post("/login", handler.Login)

	app.Get("/user", handler.UserHandlerGetAll)
	app.Get("/user/:id", middleware.Auth, handler.UserHandlerGetById)
	app.Post("/user", handler.UserHandleCreate)
	app.Put("/user/:id", handler.UserHandlerUpdateById)
	app.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	app.Delete("/user/:id", handler.UserHandlerDeleteUserById)
	app.Get("/user-post", handler.UserHandlerGetPosts)
	app.Post("/user-set-language", handler.UserSetLanguage)

	app.Get("/book", handler.BookHandlerGetAll)
	app.Post("/book", handler.BookHandleCreate)

	app.Get("/post", handler.PostHandlerGetAllPost)
	app.Post("/post", handler.PostHandlerCreate)

	app.Get("/language", handler.LanguageHandlerGetAll)
	app.Post("/language", handler.LanguageHandlerCreate)
}
