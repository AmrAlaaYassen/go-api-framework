package Visitors

import (
	"app.com/Application"
	"app.com/Models"
	tranformer "app.com/Transformers/Visitors"
	"app.com/Validations/Visitors"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	r := Application.NewRequest(c)

	// bind request with user struct
	var user Models.User
	r.Context.ShouldBind(&user)

	// validate request
	if r.ValidateRequest(Visitors.RegisterValidation(user)).Fails() {
		return
	}
	user.Token = user.Email
	user.Group = "user"
	r.DB.Create(&user)

	r.Created(tranformer.UserTransformer(user))
}

func Login(c *gin.Context) {
	r := Application.NewRequest(c)
	var user Models.User
	r.Context.ShouldBind(&user)

	if r.ValidateRequest(Visitors.LoginValidation(user)).Fails() {
		return
	}

	r.DB.Where("email = ?", user.Email).Where("password = ?", user.Password).First(&user)
	if user.ID == 0 {
		r.UserNotFound()
		return
	}
	r.Ok(tranformer.UserTransformer(user))
}
