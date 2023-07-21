package jsonschemasgenerator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/swaggest/jsonschema-go"
)

var registeredModels = map[string]interface{}{}
var jsonSchemas = map[string]jsonschema.Schema{}
var uiSchemas = map[string]UiSchema{}

func RegisterModel(name string, model interface{}) {
	registeredModels[name] = model
}

func GetSchema(name string) (interface{}, error) {
	schema, ok := jsonSchemas[name]
	if ok {
		return schema, nil
	}
	return nil, errors.New("schema not found")
}

func GetUiSchema(name string) (interface{}, error) {
	schema, ok := uiSchemas[name]
	if ok {
		return schema, nil
	}
	return nil, errors.New("schema not found")
}

func ReflectJsonSchemas() {
	reflector := jsonschema.Reflector{}
	for name, model := range registeredModels {
		schema, err := reflector.Reflect(model)
		if err != nil {
			panic(err)
		}
		jsonSchemas[name] = schema
	}
}

// ReflectUiSchemas generates the uiSchemas for the registered models, using a custom tag
// The tag is "widget" and it is used to specify the widget to use in the form
// The json tag will be used to identify the field in the json schema
func ReflectUiSchemas() {
	uiSchema := map[string]ElementType{}

	// For every model
	for name, model := range registeredModels {
		modelType := reflect.TypeOf(model)

		// For every field in the model
		for f := 0; f < modelType.NumField(); f++ {
			field := modelType.Field(f)
			// Extract the widget and json tags values
			widget := field.Tag.Get("widget")
			jsonName := field.Tag.Get("json")
			rawOptions := field.Tag.Get("options")

			// Parse the options
			optionStrings := strings.Split(rawOptions, ",")
			options := map[string]string{}
			for _, optionString := range optionStrings {
				option := strings.Split(optionString, ":")
				if len(option) == 2 {
					options[option[0]] = option[1]
				}
			}

			// If the widget and json tags are not empty
			if widget != "" && jsonName != "" {
				// Add the widget to the uiSchema
				uiSchema[jsonName] = ElementType{
					Widget:  widget,
					Options: options,
				}
			}
		}
		fmt.Printf("uiSchema for %s:\n", name)
		fmt.Println(uiSchema)
		uiSchemas[name] = uiSchema
	}
}
