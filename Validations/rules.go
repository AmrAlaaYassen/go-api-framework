package Validations

import (
	"github.com/bykovme/gotrans"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func IsRereuired() validation.Rule {
	return validation.Required.Error(gotrans.T("required"))
}

func MinMax() validation.Rule {
	return validation.Length(5, 50).Error(gotrans.T("min_max"))
}

func IsEmail() validation.Rule {
	return is.Email.Error(gotrans.T("email"))
}
