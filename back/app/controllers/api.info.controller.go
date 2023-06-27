package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// Returns all routes of the application. This route should be protected by a token and only accessible by an admin user.
func GetAllRoutes(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).
		JSON(NewResponseBody[[]fiber.Route](
			Success,
			"Server is running",
			c.App().GetRoutes(true),
		))
}
