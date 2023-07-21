package models

import "github.com/go-playground/validator/v10"

var validate = validator.New()

// ValidateStruct is a function to validate struct using its tags
// this functions can be used by any struct to validate its fields and fulfill the Validable interface
// the return value is a slice of ErrorResponse structs, which are used to return validation errors
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
