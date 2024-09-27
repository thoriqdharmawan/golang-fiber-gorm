package main

import (
	"golang-fiber-gorm/database"
	"golang-fiber-gorm/database/migration"
	"golang-fiber-gorm/route"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var corsConfig = cors.Config{
	AllowOrigins: "*",                           // Allow all origins, or specify a list of allowed origins
	AllowMethods: "GET,POST,PUT,DELETE,OPTIONS", // Allowed HTTP methods
	AllowHeaders: "Content-Type,Authorization",  // Allowed headers
}

func main() {
	database.Connect()
	migration.RunMigration()

	app := fiber.New()

	app.Use(swagger.New())
	app.Use(cors.New(corsConfig))

	route.Init(app)

	app.Listen(":8080")
}
