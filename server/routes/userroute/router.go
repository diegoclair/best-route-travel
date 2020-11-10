package userroute

import "github.com/labstack/echo"

// UserRouter holds the user handlers
type UserRouter struct {
	ctrl   *Controller
	router *echo.Echo
}

// NewRouter returns a new UserRouter instance
func NewRouter(ctrl *Controller, router *echo.Echo) *UserRouter {
	return &UserRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of user requests
func (r *UserRouter) RegisterRoutes() {
	r.router.POST("/user/login/", r.ctrl.handleSignInUser)
}
