package initializers

import (
	"example/json-schema/models"
	"os"

	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func GenerateFrontTypes() {

	path := "../front/src/types/generated"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}

	converter := typescriptify.New().
		WithInterface(true).
		WithBackupDir("").
		Add(models.UserResponse{}).
		Add(models.LogInInput{}).
		Add(models.SignUpInput{}).
		Add(models.ErrorResponse{})

	err := converter.ConvertToFile(path + "/models.ts")
	if err != nil {
		panic(err.Error())
	}

}
