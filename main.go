package main

import (
	"app.com/Application"
	"app.com/Routes"
)

func main() {

	app := Application.NewApp()

	// migrate project
	app.Migrate()
	//close app connection
	Application.CloseConnection(app)
	// routing
	routerApp := Routes.RouterApp{app}
	routerApp.Routing()

	// start server
	app.Gin.Run()
}
