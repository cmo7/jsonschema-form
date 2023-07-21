package controllers

import (
	"nartex/ngr-stack/config"
	"nartex/ngr-stack/i18n"

	"github.com/gofiber/fiber/v2"
)

/*
This file contains miscelaneus handlers, not really related to a resource directly.

TODO: Package handlers into controllers.
*/

// Returns the server configuration. This route should be protected by a token and only accessible by an admin user.
func GetServerConfiguration(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).
		JSON(NewResponseBody[fiber.Map](
			SuccessStatus,
			i18n.M(i18n.SERVER_RUNNING),
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
