package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	Gin        *gin.Engine
	DB         *gorm.DB
	Connection *sql.DB
}

func app() func() Application {
	return func() Application {
		var application Application
		application.Gin = gin.Default()
		connectToDataBase(&application)
		return application
	}
}

func (app *Application) Share() {

}
