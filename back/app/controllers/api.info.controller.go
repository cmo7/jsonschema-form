package controllers

import (
	"nartex/ngr-stack/i18n"

	"github.com/gofiber/fiber/v2"
)

// Returns all routes of the application. This route should be protected by a token and only accessible by an admin user.
func GetAllRoutes(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).
		JSON(NewResponseBody[[]fiber.Route](
			SuccessStatus,
			i18n.S(i18n.FOUND, "routes"),
			c.App().GetRoutes(true),
		))
}
