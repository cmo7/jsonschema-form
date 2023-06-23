package controllers

import (
	"example/json-schema/database"
	"example/json-schema/lib/jwthelper"
	"example/json-schema/models"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// SignInUser signs in a user. Recibes a request with a body payload in the form of a SignInInput struct
func SignUpUser(c *fiber.Ctx) error {
	var payload models.SignUpInput

	// Parse Body into payload (SignInInput struct)
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	// Validate payload using its tags using the ValidateStruct function
	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid input",
			"errors":  errors,
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
	// Hash password using bcrypt
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
	// Create new user using the SignUpInput struct and the hashed password
	newUser := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashedPassword),
		Avatar:   payload.Avatar,
	}
	// Save user in database
	result := database.DB.Create(&newUser)
	// Check for errors
	if result.Error != nil {
		switch result.Error.Error() {
		case "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)":
			// If the error is a duplicate key value, return a 400 Bad Request response
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Email already in use",
			})
		default:
			// If the error is any other, return a 502 Bad Gateway response
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

// LogInUser logs in a user. Recibes a request with a body payload in the form of a LogInInput struct
// If the user is found and the password is correct, it returns a JWT token
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
	result := database.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
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

	// Create JWT token
	tokenString, err := jwthelper.GenerateSignedToken(&user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	// Set cookie
	c.Cookie(jwthelper.GenerateTokenCookie(tokenString))

	// Return success response, including the token
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User logged in successfully",
		"data": fiber.Map{
			"user":  models.FilterUserRecord(&user),
			"token": tokenString,
		},
	})
}

// RefreshAccessToken refreshes the JWT token. This is a protected route, so it requires a valid JWT token, so it requires the ValidateToken middleware
func RefreshAccessToken(c *fiber.Ctx) error {
	if c.Locals("claims") == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
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

	// Return success response, including the token
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Token refreshed successfully",
		"data": fiber.Map{
			"user":  models.FilterUserRecord(&user),
			"token": tokenString,
		},
	})
}

// LogOutUser logs out a user. Recibes a request with a body payload in the form of a LogInInput struct
// Clears the JWT token cookie (expires it)
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

// GetCurrentUser returns the current user. This is a protected route, so it requires using the DeserializeUser middleware
func GetCurrentUser(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}
