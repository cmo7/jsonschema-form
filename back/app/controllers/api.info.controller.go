package controllers

import (
	"nartex/ngr-stack/app/types"

	"github.com/gofiber/fiber/v2"
)

// Returns all routes of the application. This route should be protected by a token and only accessible by an admin user.
func GetAllRoutes(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).
		JSON(types.NewResponseBody[[]fiber.Route](
			types.Success,
			"Server is running",
			c.App().GetRoutes(true),
		))
}
