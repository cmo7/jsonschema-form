package jwthelper

import (
	"example/json-schema/config"
	"example/json-schema/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// GenerateSignedToken generates a signed token with the user id as subject
func GenerateSignedToken(user *models.User) (string, error) {
	// Get token config
	tokenConfig := config.Options.Jwt

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": tokenConfig.Issuer,
		"sub": user.ID,
		"exp": time.Now().Add(tokenConfig.Expiration).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	})

	// Sign token with secret
	tokenString, err := token.SignedString([]byte(tokenConfig.Secret))

	return tokenString, err
}

func GenerateTokenCookie(token string) *fiber.Cookie {
	// Get token config
	tokenConfig := config.Options.Jwt

	// Create cookie
	cookie := &fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   int(tokenConfig.MaxAge.Seconds()),
		HTTPOnly: true,
		Domain:   "localhost",
	}

	return cookie
}
