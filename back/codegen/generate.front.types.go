package codegen

import (
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/config"
	"nartex/ngr-stack/utils/fronttypesgenerator"

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

	path := config.Generate.FrontTypesPath
	fronttypesgenerator.GenerateFrontTypes(path)

}
