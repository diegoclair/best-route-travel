package uploadroute

import "github.com/labstack/echo"

// UploadRouter holds the upload handlers
type UploadRouter struct {
	ctrl   *Controller
	router *echo.Echo
}

// NewRouter returns a new UploadRouter instance
func NewRouter(ctrl *Controller, router *echo.Echo) *UploadRouter {
	return &UploadRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of uploads requests
func (r *UploadRouter) RegisterRoutes() {
	r.router.POST("/prescription/:user_uuid/upload", r.ctrl.handlePrescriptionUpload)
}
