package Auth

import (
	"app.com/Application"
	"app.com/Transformers/Visitors"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	r, auth := Application.AuthRequest(c)

	if !auth {
		return
	}
	r.Ok(Visitors.UserTransformer(*r.User))
}
