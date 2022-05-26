package Visitors

import (
	"app.com/Models"
	"app.com/Validations"
	validation "github.com/go-ozzo/ozzo-validation"
)

func RegisterValidation(user Models.User) validation.Errors {
	return validation.Errors{
		"username": validation.Validate(user.Username, Validations.IsRereuired(), Validations.MinMax()),
		"email":    validation.Validate(user.Email, Validations.IsRereuired(), Validations.MinMax(), Validations.IsEmail()),
		"password": validation.Validate(user.Password, Validations.IsRereuired(), Validations.MinMax()),
	}
}
