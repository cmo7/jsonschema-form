package main

import (
	"example/json-schema/initializers"
	"example/json-schema/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	initializers.LoadEnviromentVariables()
	initializers.ConnectToDataBase()
	initializers.SyncDataBase()
}

func main() {
	// Create App
	app := fiber.New(fiber.Config{
		AppName: os.Getenv("APP_NAME"),
	})

	// Enable Middleware
	app.Use(logger.New())

	// Static Routes
	app.Static("/", "./public").Name("public")

	// Add Routes
	routes.AddRoutes(app)

	// Run Server
	app.Listen(":" + os.Getenv("PORT"))

}
