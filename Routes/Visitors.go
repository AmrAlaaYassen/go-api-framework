package Routes

import "app.com/Controllers/Visitors"

func (app RouterApp) visitorRoutes() {
	app.Gin.GET("/createuser", Visitors.CreateUser)
}
