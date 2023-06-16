package controllers

import (
	"example/json-schema/initializers"
	"example/json-schema/models"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(c *fiber.Ctx) error {
	var payload models.SignUpInput

	// Parse Body
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	// Validate payload using its tags
	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": errors,
		})
	}
	// Check if password and password confirmation are equal
	// This must be also done in the frontend
	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Password and Password Confirmation must be equal",
		})
	}
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(payload.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	// Create new user
	newUser := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashedPassword),
		Avatar:   payload.Avatar,
	}
	// Save user in database
	result := initializers.DB.Create(&newUser)
	// Check for errors
	if result.Error != nil {
		switch result.Error.Error() {
		case "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)":
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Email already in use",
			})
		default:
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"status":  "error",
				"message": result.Error.Error(),
			})
		}
	}
	// Return success response, including the new user (filtered)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User created successfully",
		"data":    models.FilterUserRecord(&newUser),
	})

}

func LogInUser(c *fiber.Ctx) error {
	var payload models.LogInInput

	// Parse Body
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	// Validate payload using its tags
	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": errors,
		})
	}
	// Find user by email
	var user models.User
	result := initializers.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	// Check for errors
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
		})
	}
	// Check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid credentials",
		})
	}
	// Return success response, including the user (filtered)

	tokenByte := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = user.ID
	claims["exp"] = now.Add(time.Hour).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	// Set cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   3600,
		HTTPOnly: true,
		Domain:   "localhost",
	})
	// Return success response, including the token
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User logged in successfully",
		"token":   tokenString,
	})
}

func LogOutUser(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24) // Set cookie expiration date to yesterday
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User logged out successfully",
	})
}

func GetCurrentUser(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": user,
		},
	})
}
