package controllers

import (
	"example/json-schema/config"

	"github.com/gofiber/fiber/v2"
)

// Returns the server configuration. This route should be protected by a token and only accessible by an admin user.
func GetServerConfiguration(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Server is running",
		"data":    config.Options,
	})
}

// Returns all routes of the application. This route should be protected by a token and only accessible by an admin user.
func GetAllRoutes(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Server is running",
		"data":    c.App().GetRoutes(true),
	})
}
