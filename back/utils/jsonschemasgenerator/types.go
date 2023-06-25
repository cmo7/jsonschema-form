package jsonschemasgenerator

type ElementType struct {
	Widget  string            `json:"ui:widget,omitempty"`
	Options map[string]string `json:"ui:options,omitempty"`
}

type UiSchema map[string]ElementType
