package Routes

import "app.com/Controllers/Visitors"

func (app RouterApp) visitorRoutes() {
	app.Gin.POST("register", Visitors.Register)
}
