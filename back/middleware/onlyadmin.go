package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func OnlyAdmin(c *fiber.Ctx) error {

	if c.Locals("claims") == nil {
		log.Printf("no Claims")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized (no claims)",
		})
	}

	claims := c.Locals("claims").(jwt.MapClaims)
	if claims["role"] != "admin" {
		log.Printf("Role: %s", claims["role"])
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized (not admin)",
		})
	}
	return c.Next()
}
