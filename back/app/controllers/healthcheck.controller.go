package controllers

import (
	"nartex/ngr-stack/i18n"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": i18n.M(i18n.HEALTH_CHECK),
	})
}
