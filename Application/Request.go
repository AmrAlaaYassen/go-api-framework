package Application

import (
	"app.com/Models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SharedResources interface {
	Share()
}

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
	User       *Models.User
	Auth       bool
}

func (req *Request) Share() {}

func req() func(c *gin.Context) Request {
	return func(c *gin.Context) Request {
		var request Request
		request.Context = c
		return request
	}
}

// init new request
func NewRequest(c *gin.Context) Request {
	request := req()
	req := request(c)
	connectToDataBase(&req)
	return req
}

// Response standard
func (req Request) Response(code int, body interface{}) {
	CloseConnection(&req)
	req.Context.JSON(code, body)
}
