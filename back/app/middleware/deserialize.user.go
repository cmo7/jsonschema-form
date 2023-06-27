package middleware

import (
	"fmt"
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/database"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DeserializeUser(c *fiber.Ctx) error {

	// Get claims from context
	claims := c.Locals("claims").(jwt.MapClaims)
	// Find user
	var user models.User
	database.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	// If user is not present
	if user.ID.String() == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	// Set user in locals
	c.Locals("user", models.FilterUserRecord(&user))

	// Continue
	return c.Next()
}
