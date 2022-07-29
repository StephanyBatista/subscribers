package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validate(body interface{}, c *gin.Context) bool {

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make([]string, len(validationErrors))
		for index, fieldError := range validationErrors {
			message := "'" + fieldError.Field() + "'"
			if fieldError.Tag() == "required" {
				message += " is required"
			} else if fieldError.Tag() == "email" {
				message += " is invalid"
			}

			errors[index] = message
		}
		c.JSON(http.StatusBadRequest, NewErrorsReponse(errors))
		return false
	}
	return true
}
