package Routes

import "app.com/Application"

type RouterApp struct {
	*Application.Application
}

func (router RouterApp) Routing() {
	router.authRoutes()
	router.visitorRoutes()

}
