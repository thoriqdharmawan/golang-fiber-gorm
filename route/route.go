package route

import (
	"golang-fiber-gorm/handler"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Get("/user", handler.UserHandlerGetAll)
	app.Get("/user/:id", handler.UserHandlerGetById)
	app.Post("/user", handler.UserHandleCreate)
}
