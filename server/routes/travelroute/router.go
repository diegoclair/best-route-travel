package travelroute

import "github.com/labstack/echo"

// TravelRouter holds the user handlers
type TravelRouter struct {
	ctrl   *Controller
	router *echo.Echo
}

// NewRouter returns a new TravelRouter instance
func NewRouter(ctrl *Controller, router *echo.Echo) *TravelRouter {
	return &TravelRouter{
		ctrl:   ctrl,
		router: router,
	}
}

// RegisterRoutes is a routers map of travel requests
func (r *TravelRouter) RegisterRoutes() {
	r.router.GET("/travel/:where_from/:where_to/bestroute/", r.ctrl.handleGetTravelBestRoute)
	r.router.POST("/travel/", r.ctrl.handleAddNewRoute)
}
