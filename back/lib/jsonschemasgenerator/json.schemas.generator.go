package jsonschemasgenerator

import (
	"errors"

	"github.com/swaggest/jsonschema-go"
)

var registeredModels = map[string]interface{}{}
var jsonSchemas = map[string]jsonschema.Schema{}

func RegisterModel(name string, model interface{}) {
	registeredModels[name] = model
}

func GetSchema(name string) (interface{}, error) {
	schema, ok := jsonSchemas[name]
	if ok {
		return schema, nil
	}
	return nil, errors.New("Schema not found")
}

func ReflectSchemas() {
	reflector := jsonschema.Reflector{}
	for name, model := range registeredModels {
		schema, err := reflector.Reflect(model)
		if err != nil {
			panic(err)
		}
		jsonSchemas[name] = schema
	}
}
