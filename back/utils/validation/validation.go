package validation

import "github.com/go-playground/validator/v10"

type Validable interface {
	Validate() []*ErrorResponse
}

// ErrorResponse is the response payload for error response
type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value,omitempty"`
}

var validate = validator.New()

// ValidateStruct is a function to validate struct using its tags
func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorResponse{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}
