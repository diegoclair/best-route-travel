package server

import (
	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/best-route-travel/server/routes/pingroute"
	"github.com/diegoclair/best-route-travel/server/routes/travelroute"
	"github.com/diegoclair/best-route-travel/service"
	"github.com/labstack/echo"
)

type controller struct {
	pingController   *pingroute.Controller
	travelController *travelroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service, svm service.Manager) *echo.Echo {
	mapper := mapper.New()

	srv := echo.New()

	travelService := svm.TravelService(svc)

	return setupRoutes(srv, &controller{
		pingController:   pingroute.NewController(),
		travelController: travelroute.NewController(travelService, mapper),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *echo.Echo, s *controller) *echo.Echo {

	pingroute.NewRouter(s.pingController, srv).RegisterRoutes()
	travelroute.NewRouter(s.travelController, srv).RegisterRoutes()

	return srv
}
