package main

import (
	"example/json-schema/initializers"
	"example/json-schema/routes"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
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
	initializers.LoadConfig(development)
	initializers.ConnectToDataBase()
	initializers.InitializeApp()

	generate := initializers.Config.Generate
	if generate.AutoMigrate {
		initializers.SyncDataBase()
	}
	if generate.FrontTypes {
		initializers.GenerateFrontTypes()
	}
	initializers.GenerateJsonFormSchemas()
}

func main() {
	config := initializers.Config
	// Create App
	app := initializers.App

	// Register App Middleware
	if config.Logger.MainLogger {
		app.Use(logger.New())
	}
	if config.Middleware.Helmet {
		app.Use(helmet.New())
	}
	if config.Middleware.Compress {
		app.Use(compress.New())
	}
	if config.Middleware.Cache {
		app.Use(cache.New())
	}

	app.Use(cors.New(cors.Config{
		AllowHeaders: strings.Join([]string{
			"Origin",
			"Content-Type",
			"Accept",
			"Access-Control-Allow-Headers",
			"Authorization",
			"X-Requested-With",
		}, ","),
		AllowOrigins: strings.Join([]string{
			config.Client.URL,
		}, ","),
	}))

	// Client Serving Mode
	switch config.Client.Mode {
	case "internal":
		app.Static("/", "./public").Name("Serving Client (Internal)")
	case "external":
		app.Get("/", func(c *fiber.Ctx) error {
			return c.Redirect(config.Client.URL)
		}).Name("Client Redirect (External)")
	default:
		panic("CLIENT_MODE not defined or invalid")
	}

	// Add Routes
	routes.AddApiRoutes(app)

	// Run Server
	webConfig := config.WebServer
	portString := fmt.Sprintf(":%d", webConfig.Port)

	if webConfig.TLS {
		err := app.ListenTLS(portString, webConfig.CertFile, webConfig.KeyFile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := app.Listen(portString)
		if err != nil {
			log.Fatal(err)
		}
	}

}
