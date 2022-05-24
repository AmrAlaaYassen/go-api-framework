package Visitors

import (
	"app.com/Application"
	"app.com/Models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	r := Application.NewRequest(c)
	user := Models.User{
		Username: "amr",
		Email:    "email",
		Password: "password",
	}

	r.DB.Create(&user)
	r.Created(user)
}
