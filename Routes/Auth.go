package Routes

import "app.com/Controllers/Auth"

func (app RouterApp) authRoutes() {
	app.Gin.GET("/createuser", Auth.CreateUser)
}
