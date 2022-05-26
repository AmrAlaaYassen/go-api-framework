package Visitors

import (
	"app.com/Application"
	"app.com/Models"
	"app.com/Validations/Visitors"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	r := Application.NewRequest(c)
	// bind request with user struct
	var user Models.User
	r.Context.ShouldBind(&user)
	// validate request
	r.ValidateRequest(Visitors.RegisterValidation(user))

	if r.ValidationError != nil {
		r.BadRequest(r.ValidationError)
		return
	}

	r.Ok(user)
}
