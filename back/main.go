package main

import (
	"example/json-schema/initializers"
	"example/json-schema/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://localhost:4173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Static Routes
	app.Static("/", "./public").Name("public")

	// Add Routes
	routes.AddRoutes(app)

	// Run Server
	err := app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}

}
