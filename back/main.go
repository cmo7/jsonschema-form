package main

import (
	"example/json-schema/initializers"
	"example/json-schema/routes"
	"fmt"
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
	// Create App
	app := fiber.New(fiber.Config{
		AppName: initializers.Config.AppName,
	})

	// Register Middleware
	if initializers.Config.Logger.MainLogger {
		app.Use(logger.New())
	}

	midlewareConfig := initializers.Config.Middleware

	if midlewareConfig.Helmet {
		app.Use(helmet.New())
	}

	if midlewareConfig.Compress {
		app.Use(compress.New())
	}

	if midlewareConfig.Cache {
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
			initializers.Config.Client.URL,
		}, ","),
	}))

	// Client Serving Mode
	switch initializers.Config.Client.Mode {
	case "internal":
		app.Static("/", "./public").Name("public")
	case "external":
		app.Get("/", func(c *fiber.Ctx) error {
			return c.Redirect(initializers.Config.Client.URL)
		})
	default:
		panic("CLIENT_MODE not defined or invalid")
	}

	// Static Routes

	// Add Routes
	routes.AddRoutes(app)

	// Run Server

	webConfig := initializers.Config.WebServer
	portString := fmt.Sprintf(":%d", webConfig.Port)

	if initializers.Config.WebServer.TLS {
		err := app.ListenTLS(portString, webConfig.CertFile, webConfig.KeyFile)
		if err != nil {
			panic(err)
		}
	} else {
		err := app.Listen(portString)
		if err != nil {
			panic(err)
		}
	}

}
