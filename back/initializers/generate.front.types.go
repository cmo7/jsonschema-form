package initializers

import (
	"example/json-schema/config"
	"example/json-schema/lib/fronttypesgenerator"
	"example/json-schema/models"

	"github.com/gofiber/fiber/v2"
)

// GenerateFrontTypes generates typescript models from Go models
func GenerateFrontTypes() {
	// Add models here
	fronttypesgenerator.RegisterModel(models.UserResponse{})
	fronttypesgenerator.RegisterModel(models.LogInInput{})
	fronttypesgenerator.RegisterModel(models.SignUpInput{})
	fronttypesgenerator.RegisterModel(models.ErrorResponse{})
	fronttypesgenerator.RegisterModel(fiber.Route{})

	path := config.Options.Generate.FrontTypesPath
	fronttypesgenerator.GenerateFrontTypes(path)

}
