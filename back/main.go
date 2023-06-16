package main

import (
	"example/json-schema/initializers"
	"example/json-schema/routes"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Enviroments
const (
	development   = "development"
	preproduction = "preproduction"
	production    = "production"
	container     = "container"
)

func init() {
	initializers.LoadEnviromentVariables(development)
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
		AllowHeaders: "Origin, Content-Type, Accept",
		// Add AllowOrigins to .env
		AllowOrigins: strings.Join([]string{
			os.Getenv("CLIENT_URL"),
		}, ","),
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
