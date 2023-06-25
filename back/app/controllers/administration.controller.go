package controllers

import (
	"nartex/ngr-stack/app/types"
	"nartex/ngr-stack/config"

	"github.com/gofiber/fiber/v2"
)

// Returns the server configuration. This route should be protected by a token and only accessible by an admin user.
func GetServerConfiguration(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).
		JSON(types.NewResponseBody[fiber.Map](
			types.Success,
			"Server is running",
			fiber.Map{
				"App":        config.App,
				"Database":   config.Database,
				"Debug":      config.Debug,
				"Generate":   config.Generate,
				"Middleware": config.Middleware,
				"WebServer":  config.WebServer,
				"Client":     config.Client,
				"Jwt":        config.Jwt,
				"Logger":     config.Logger,
				"Pagination": config.Pagination,
			}))
}
