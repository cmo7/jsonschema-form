package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func MainPage(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}
