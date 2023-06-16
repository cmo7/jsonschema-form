package models

import "github.com/go-playground/validator/v10"

// validate is a validator instance for validating struct using its tags
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
