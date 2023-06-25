package app

import (
	"fmt"
	"nartex/ngr-stack/config"
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

func BootstrapApp() *fiber.App {

	appNameString := fmt.Sprintf("%s (%s)", config.App.AppName, config.App.Enviroment)
	app := fiber.New(fiber.Config{
		AppName: appNameString,
	})

	fmt.Printf("App Name: %s\n", appNameString)

	// No panics, just recover
	if config.App.Enviroment == config.Production {
		app.Use(recover.New())
	}

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

	// Client Serving Mode
	switch config.Client.Mode {
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
			config.Client.URL,
		}, ", ")

		app.Use(cors.New(cors.Config{
			AllowHeaders: allowHeaders,
			AllowOrigins: allowOrigins,
		}))
		// Redirect to external client
		app.Get("/", func(c *fiber.Ctx) error {
			return c.Redirect(config.Client.URL)
		}).Name("Client Redirect (External)")
	default:
		panic("CLIENT_MODE not defined or invalid")
	}
	// Add DevTools
	if config.Debug.DevTools {
		app.Get("/metrics", monitor.New()).Name("Fiber Metrics")
	}

	return app
}
