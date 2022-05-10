package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := app()
	server := app()
	server.Gin.GET("/ping", func(context *gin.Context) {
		req := newRequest(context)
		req.ok("Thank you...")

	})
	server.Gin.Run()
}
