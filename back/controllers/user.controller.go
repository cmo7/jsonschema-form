package controllers

import (
	"example/json-schema/initializers"
	"example/json-schema/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/swaggest/jsonschema-go"
)

func UserSchema(c *fiber.Ctx) error {
	schemaName := c.Params("schemaName")
	reflector := jsonschema.Reflector{}

	// Get correct Model
	var model interface{}
	switch schemaName {
	case "SiginInput":
		model = models.SignUpInput{}
	case "LogInInput":
		model = models.LogInInput{}
	case "UserResponse":
		model = models.UserResponse{}
	case "ErrorResponse":
		model = models.ErrorResponse{}
	default:
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fiber.ErrNotFound.Message,
			"data":    nil,
		})
	}

	// Reflect and teg schema
	schema, err := reflector.Reflect(model)
	log.Println(schema)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fiber.ErrInternalServerError.Message,
			"data":    err,
		})
	}
	return c.JSON(schema)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id is required",
			"data":    nil,
		})
	}
	var user = new(models.User)
	initializers.DB.First(&user, "id = ?", id)
	if user.ID == nil || user.ID.String() == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fiber.ErrNotFound.Message,
			"status":  "error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.FilterUserRecord(user))
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	initializers.DB.Find(&users)
	var usersResponse []models.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, models.FilterUserRecord(&user))
	}
	return c.Status(fiber.StatusOK).JSON(usersResponse)
}
