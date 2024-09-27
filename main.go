package main

import (
	"golang-fiber-gorm/database"
	"golang-fiber-gorm/database/migration"
	"golang-fiber-gorm/route"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	migration.RunMigration()

	app := fiber.New()
	app.Use(swagger.New())

	route.Init(app)

	app.Listen(":8080")
}
