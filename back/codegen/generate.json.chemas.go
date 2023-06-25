package codegen

import (
	"nartex/ngr-stack/app/models"
	"nartex/ngr-stack/utils/jsonschemasgenerator"
)

// GenerateJsonFormSchemas generates json schemas from Go models
func GenerateJsonFormSchemas() {
	// Register models here. Use distinct names for each model
	// The name will be used as the key for the schema
	// The value will be the model itself
	// Keys are used to retrieve the schema later
	jsonschemasgenerator.RegisterModel("UserResponse", models.UserResponse{})
	jsonschemasgenerator.RegisterModel("LogInInput", models.LogInInput{})
	jsonschemasgenerator.RegisterModel("SignUpInput", models.SignUpInput{})

	// Reflect all registered models. Call this function after all models are registered
	jsonschemasgenerator.ReflectJsonSchemas()
	jsonschemasgenerator.ReflectUiSchemas()
}
