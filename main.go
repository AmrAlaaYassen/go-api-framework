package main

import (
	"app.com/Application"
	"app.com/Models"
	"app.com/Routes"
)

func main() {

	app := Application.NewApp()

	// migrate project
	app.DB.AutoMigrate(&Models.User{})
	//close app connection
	Application.CloseConnection(app)
	// routing
	routerApp := Routes.RouterApp{app}
	routerApp.Routing()

	// start server
	app.Gin.Run()
}
