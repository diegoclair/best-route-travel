package server

import (
	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/best-route-travel/server/routes/pingroute"
	"github.com/diegoclair/best-route-travel/server/routes/uploadroute"
	"github.com/diegoclair/best-route-travel/server/routes/userroute"
	"github.com/diegoclair/best-route-travel/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type controller struct {
	pingController   *pingroute.Controller
	userController   *userroute.Controller
	uploadController *uploadroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service, svm service.Manager) *echo.Echo {
	mapper := mapper.New()

	srv := echo.New()

	userService := svm.UserService(svc)
	uploadService := svm.UploadService(svc)

	//CORS
	srv.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	return setupRoutes(srv, &controller{
		pingController:   pingroute.NewController(),
		userController:   userroute.NewController(userService, mapper),
		uploadController: uploadroute.NewController(uploadService),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *echo.Echo, s *controller) *echo.Echo {

	pingroute.NewRouter(s.pingController, srv).RegisterRoutes()
	userroute.NewRouter(s.userController, srv).RegisterRoutes()
	uploadroute.NewRouter(s.uploadController, srv).RegisterRoutes()

	return srv
}
