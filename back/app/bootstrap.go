package app

import (
	"example/json-schema/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func BootstrapApp(options *config.ConfigOptions) *fiber.App {

	app := fiber.New()

	// Initialize default config

	// No panics, just recover
	app.Use(recover.New())

	// Register App Middleware
	if options.Logger.MainLogger {
		app.Use(logger.New())
	}
	if options.Middleware.Helmet {
		app.Use(helmet.New())
	}
	if options.Middleware.Compress {
		app.Use(compress.New())
	}
	if options.Middleware.Cache {
		app.Use(cache.New())
	}

	// Client Serving Mode
	switch options.Client.Mode {
	case config.Internal:
		// Serve internal client from public folder
		app.Static("/", "./public").Name("Serving Client (Internal)")
	case config.External:
		// CORS Middleware for external client

		allowHeaders := strings.Join([]string{
			"Origin",
			"Content-Type",
			"Accept",
			"Access-Control-Allow-Headers",
			"Authorization",
			"X-Requested-With",
		}, ", ")

		allowOrigins := strings.Join([]string{
			options.Client.URL,
		}, ", ")

		app.Use(cors.New(cors.Config{
			AllowHeaders: allowHeaders,
			AllowOrigins: allowOrigins,
		}))
		// Redirect to external client
		app.Get("/", func(c *fiber.Ctx) error {
			return c.Redirect(options.Client.URL)
		}).Name("Client Redirect (External)")
	default:
		panic("CLIENT_MODE not defined or invalid")
	}
	// Add DevTools
	if options.DevTools {
		app.Get("/metrics", monitor.New()).Name("Fiber Metrics")
	}

	return app
}
