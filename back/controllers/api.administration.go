package controllers

import (
	"example/json-schema/initializers"

	"github.com/gofiber/fiber/v2"
)

func GetServerConfiguration(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Server is running",
		"data":    initializers.Config,
	})
}

func GetAllRoutes(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Server is running",
		"data":    initializers.App.GetRoutes(true),
	})
}
