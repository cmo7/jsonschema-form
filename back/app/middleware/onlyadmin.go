package middleware

import (
	"fmt"
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/database"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// OnlyAdmin is a middleware that checks if the user is an admin before continuing the request
func OnlyAdmin(c *fiber.Ctx) error {

	if c.Locals("claims") == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized (no claims)",
		})
	}

	claims := c.Locals("claims").(jwt.MapClaims)
	var user models.User
	result := database.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	for _, role := range user.Roles {
		if role.Name == "admin" {
			return c.Next()
		}
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  "error",
		"message": "Unauthorized (not admin)",
	})

}
