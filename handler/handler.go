package handler

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type (
	Handler struct {
		DB        *gorm.DB
		Validator CustomValidator
	}
	CustomValidator struct {
		V *validator.Validate
	}
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)

func (cv *CustomValidator) Validate(i interface{}) map[string]string {
	errMap := make(map[string]string)

	if err := cv.V.Struct(i); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, validationErr := range validationErrors {
				field := validationErr.Field() // Get the field name
				message := validationErr.Tag() // Get the error message
				errMap[field] = message
			}
		}
		return errMap

	}
	return nil

}
