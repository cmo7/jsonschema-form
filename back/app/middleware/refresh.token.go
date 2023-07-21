package middleware

import (
	"fmt"
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/database"
	"nartex/ngr-stack/utils/jwthelper"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// RefreshAccessToken is a middleware that refreshes the access token and sets it in the response header
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
