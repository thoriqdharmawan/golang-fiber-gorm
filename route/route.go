package route

import (
	"golang-fiber-gorm/handler"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Get("/user", handler.UserHandlerGetAll)
	app.Get("/user/:id", handler.UserHandlerGetById)
	app.Post("/user", handler.UserHandleCreate)
	app.Put("/user/:id", handler.UserHandlerUpdateById)
	app.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	app.Delete("/user/:id", handler.UserHandlerDeleteUserById)
}
