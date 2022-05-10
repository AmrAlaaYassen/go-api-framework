package main

import (
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
func newRequest(c *gin.Context) Request {
	request := req()
	req := request(c)
	connectToDataBase(&req)
	return req
}

func (req Request) ok(body interface{}) {
	req.Response(200, body)
}

// Response standard
func (req Request) Response(code int, body interface{}) {
	closeConnection(&req)
	req.Context.JSON(code, body)
}
