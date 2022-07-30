package domain

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func Validate(body interface{}) []error {

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errs := make([]error, len(validationErrors))
		for index, fieldError := range validationErrors {
			message := "'" + fieldError.Field() + "'"
			if fieldError.Tag() == "required" {
				message += " is required"
			} else if fieldError.Tag() == "email" {
				message += " is invalid"
			}

			errs[index] = errors.New(message)
		}

		return errs
	}
	return nil
}
