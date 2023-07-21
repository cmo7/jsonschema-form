package middleware

import (
	"fmt"
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/database"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// DeserializeUser is a middleware that deserializes the user from the JWT and sets it in locals
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
	c.Locals("user", user.ToDto())

	// Continue
	return c.Next()
}
