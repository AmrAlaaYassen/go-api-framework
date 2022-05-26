package main

import (
	"app.com/Application"
	"app.com/Routes"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load(".env")
	app := Application.NewApp()

	// migrate project
	app.Migrate()
	// seed DB
	app.Seed()
	//close app connection
	Application.CloseConnection(app)
	// routing
	routerApp := Routes.RouterApp{app}
	routerApp.Routing()

	// start server
	app.Gin.Run()
}
