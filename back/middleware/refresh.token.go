package middleware

import (
	"example/json-schema/database"
	"example/json-schema/lib/jwthelper"
	"example/json-schema/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func RefreshAccessToken(c *fiber.Ctx) error {

	// Generate new token
	claims := c.Locals("claims").(jwt.MapClaims)
	var user models.User
	result := database.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	// Create new JWT token
	tokenString, err := jwthelper.GenerateSignedToken(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	// Set cookie
	c.Cookie(jwthelper.GenerateTokenCookie(tokenString))

	// Add token to response header
	c.Set("X-Access-Token", tokenString)
	// Deserialize response body
	return c.Next()
}
