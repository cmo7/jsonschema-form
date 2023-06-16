package middleware

import (
	"example/json-schema/initializers"
	"example/json-schema/models"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DeserializeUser(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.Split(authorization, " ")[1]
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

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

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	var user models.User
	initializers.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	if user.ID == nil || user.ID.String() == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	c.Locals("user", models.FilterUserRecord(&user))

	return c.Next()
}
