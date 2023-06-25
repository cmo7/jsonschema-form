package fronttypesgenerator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

var registeredModels = []interface{}{}

func RegisterModel(model interface{}) {
	registeredModels = append(registeredModels, model)
}

func GenerateFrontTypes(path string) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}

	converter := typescriptify.New().
		WithInterface(true).
		WithBackupDir("")

	for _, model := range registeredModels {
		converter.Add(model)
	}

	fmt.Println("Generating front types...")
	absFilePath, _ := filepath.Abs(path)
	fmt.Println("Path: " + absFilePath)
	err := converter.ConvertToFile(path + "/models.ts")
	if err != nil {
		panic(err.Error())
	}
}
