package main

import (
	"example/json-schema/initializers"
	"example/json-schema/routes"
	"os"
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
	initializers.LoadEnviromentVariables(development)
	initializers.ConnectToDataBase()
	if os.Getenv("DB_AUTO_MIGRATE") == "true" {
		initializers.SyncDataBase()
	}
	if os.Getenv("GENERATE_INTERFACES") == "true" {
		initializers.GenerateFrontTypes()
	}
}

func main() {
	// Create App
	app := fiber.New(fiber.Config{
		AppName: os.Getenv("APP_NAME"),
	})

	// Register Middleware

	if os.Getenv("LOG_ENABLED") == "true" || os.Getenv("LOG_ENABLED") == "" {
		app.Use(logger.New())
	}

	if os.Getenv("MID_HELMET") == "true" || os.Getenv("MID_HELMET") == "" {
		app.Use(helmet.New())
	}

	if os.Getenv("MID_COMPRESS") == "true" || os.Getenv("MID_COMPRESS") == "" {
		app.Use(compress.New())
	}

	if os.Getenv("MID_CACHE") == "true" || os.Getenv("MID_CACHE") == "" {
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
			os.Getenv("CLIENT_URL"),
		}, ","),
	}))

	// Client Serving Mode
	clientMode := os.Getenv("CLIENT_MODE")
	switch clientMode {
	case "internal":
		app.Static("/", "./public").Name("public")
	case "external":
		app.Get("/", func(c *fiber.Ctx) error {
			return c.Redirect(os.Getenv("CLIENT_URL"))
		})
	default:
		panic("CLIENT_MODE not defined or invalid")
	}

	// Static Routes

	// Add Routes
	routes.AddRoutes(app)

	// Run Server

	if os.Getenv("TLS_ENABLED") == "true" {
		err := app.ListenTLS(":"+os.Getenv("TLS_PORT"), os.Getenv("TLS_CERT"), os.Getenv("TLS_KEY"))
		if err != nil {
			panic(err)
		}
	} else {
		err := app.Listen(":" + os.Getenv("PORT"))
		if err != nil {
			panic(err)
		}
	}

}
