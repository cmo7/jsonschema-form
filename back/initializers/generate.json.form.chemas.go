package initializers

import (
	"example/json-schema/lib"
	"example/json-schema/models"
)

func GenerateJsonFormSchemas() {
	// Register models here. Use distinct names for each model
	// The name will be used as the key for the schema
	// The value will be the model itself
	// Keys are used to retrieve the schema later
	lib.RegisterModel("UserResponse", models.UserResponse{})
	lib.RegisterModel("LogInInput", models.LogInInput{})
	lib.RegisterModel("SignUpInput", models.SignUpInput{})

	// Reflect all registered models. Call this function after all models are registered
	lib.ReflectSchemas()
}
