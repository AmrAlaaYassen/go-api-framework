package Application

import (
	"app.com/Database"
	"app.com/Database/Seeders"
	"database/sql"
	"fmt"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	Gin        *gin.Engine
	DB         *gorm.DB
	Connection *sql.DB
}

func App() func() *Application {
	return func() *Application {
		var application Application
		application.Gin = gin.Default()
		connectToDataBase(&application)
		err := gotrans.InitLocales("./public/lang")

		if err != nil {
			fmt.Println("Error on loading trans files ..", err.Error())
		}
		return &application
	}
}

func (app *Application) Share() {

}

func NewApp() *Application {
	app := App()
	return app()
}

func (app *Application) Migrate() {
	Database.Migrate(app.DB)
}

func (app *Application) Seed() {
	Seeders.Seed(app.DB)
}
