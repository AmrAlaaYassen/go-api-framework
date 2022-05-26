package Application

import (
	"app.com/Models"
	"database/sql"
	"fmt"
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
	IsAuth     bool
	IsAdmin    bool
}

func (req *Request) Share() {}

func req() func(c *gin.Context) *Request {
	return func(c *gin.Context) *Request {
		var request Request
		request.Context = c
		connectToDataBase(&request)
		return &request
	}
}

// init new request
func NewRequest(c *gin.Context) *Request {
	request := req()
	return request(c)
}

func NewRequestWithAuth(c *gin.Context) *Request {
	return NewRequest(c).Auth()
}

// Response standard
func (req Request) Response(code int, body interface{}) {
	CloseConnection(&req)
	req.Context.JSON(code, body)
}

func AuthRequest(c *gin.Context) (*Request, bool) {
	r := NewRequestWithAuth(c)
	if !r.IsAuth {
		r.NotAuth()
		return r, false
	}

	return r, true
}

func (req *Request) Auth() *Request {
	req.IsAuth = false
	req.IsAdmin = false
	authHeader := req.Context.GetHeader("Authorization")

	if authHeader != "" {
		fmt.Println(authHeader)
		req.DB.Where("token = ? ", authHeader).First(&req.User)
		//req.DB.Find(&req.User, "token = ?", authHeader)
		if req.User.ID != 0 {
			req.IsAuth = true
			if req.User.Group == "admin" {
				fmt.Println("inside admin")
				req.IsAdmin = true
			}
		}
	}
	return req
}
