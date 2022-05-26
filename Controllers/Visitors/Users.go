package Visitors

import (
	"app.com/Application"
	"app.com/Models"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func Register(c *gin.Context) {
	r := Application.NewRequestWithAuth(c)
	// bind request with user struct
	var user Models.User
	r.Context.ShouldBind(&user)
	// validate request
	err := validation.Errors{
		"username": validation.Validate(user.Username, validation.Required.Error(gotrans.T("required")), validation.Length(5, 50).Error(gotrans.T("min_max"))),
		"email":    validation.Validate(user.Email, validation.Required.Error(gotrans.T("required")), is.Email.Error(gotrans.T("email")), validation.Length(5, 50).Error(gotrans.T("min_max"))),
		"password": validation.Validate(user.Password, validation.Required.Error(gotrans.T("required")), validation.Length(5, 50).Error(gotrans.T("min_max"))),
	}.Filter()

	if err != nil {
		r.BadRequest(err)
		return
	}

	r.Ok(user)
}
