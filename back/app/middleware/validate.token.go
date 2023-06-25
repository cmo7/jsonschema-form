package middleware

import (
	"fmt"
	"nartex/ngr-stack/config"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// ValidateToken middleware validates JWT token and adds its claims to context
func ValidateToken(c *fiber.Ctx) error {
	var tokenString string

	jwtConfig := config.Jwt

	// Parse Authorization header
	authorization := c.Get("Authorization")
	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.Split(authorization, " ")[1]
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	// If token is not present
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "No token provided",
		})
	}

	// Parse token
	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	// Validate token
	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	if claims["sub"] == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	if claims["iss"] == nil || claims["iss"].(string) != jwtConfig.Issuer {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	// Add claims to context
	c.Locals("claims", claims)

	// continue stack
	return c.Next()
}
