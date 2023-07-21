package app

import (
	"fmt"
	"nartex/ngr-stack/app/middleware"
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

	// Create the App, using the Appname and Enviroment from the config.
	appName := fmt.Sprintf("%s (%s)", config.App.AppName, config.App.Enviroment)
	app := fiber.New(fiber.Config{
		AppName: appName,
	})

	// Register App Middleware
	//----------------------------------------------------------------------------

	app.Use(middleware.RequestTimer)

	// In case of panic, recover
	if config.App.Enviroment == config.Production {
		app.Use(recover.New())
	}

	// Activate logger, with specified configuration
	if config.Logger.MainLogger {
		app.Use(logger.New())
	}

	// Activate Helmet, setting some standard HTTP headers for extra security
	if config.Middleware.Helmet {
		app.Use(helmet.New())
	}

	// Compresses responses with gzip, deflate and brotli
	if config.Middleware.Compress {
		app.Use(compress.New())
	}

	// Caches responses
	if config.Middleware.Cache {
		app.Use(cache.New())
	}

	// Specifies how will we be serving our client app
	//----------------------------------------------------------------------------
	// Client Serving Mode
	switch config.Client.Mode {
	case config.Internal:
		// Serve internal client from public folder
		app.Static("/", "./public").Name("Serving Client (Internal)")
	case config.External:
		// CORS Middleware for external client
		// As the request will be coming from another origin, we need to enable it
		allowHeaders := strings.Join([]string{
			"Origin",
			"Content-Type",
			"Accept",
			"Access-Control-Allow-Headers",
			"Authorization",
			"X-Requested-With",
		}, ", ")

		// Enabled origins are at least the client url.
		allowOrigins := strings.Join([]string{
			config.Client.URL,
		}, ", ")

		app.Use(cors.New(cors.Config{
			AllowHeaders: allowHeaders,
			AllowOrigins: allowOrigins,
		}))

		// Redirect to external client
		// We wont need CORS in this case, at least for the client.
		app.Get("/", func(c *fiber.Ctx) error {
			return c.Redirect(config.Client.URL)
		}).Name("Client Redirect (External)")
	default:
		panic("CLIENT_MODE not defined or invalid")
	}

	// Add DevTools
	// Monitor allows us to profile the resource consumption of the app
	if config.Debug.DevTools {
		app.Get("/metrics", monitor.New()).Name("Fiber Metrics")
	}

	// Returns the configured Fiber App.
	return app
}
