package Auth

import (
	"app.com/Application"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	r, auth := Application.AuthRequest(c)

	if !auth {
		return
	}
	r.Ok(gotrans.T("hello"))
}
