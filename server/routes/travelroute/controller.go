package travelroute

import (
	"net/http"
	"strings"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/best-route-travel/domain/contract"
	"github.com/diegoclair/best-route-travel/domain/entity"
	"github.com/diegoclair/best-route-travel/server/viewmodel"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/labstack/echo"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller is a interface to interact with services
type Controller struct {
	travelService contract.TravelService
	mapper        mapper.Mapper
}

//NewController to handle requests
func NewController(travelService contract.TravelService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			travelService: travelService,
			mapper:        mapper,
		}
	})
	return instance
}

func (c Controller) handleGetTravelBestRoute(ctx echo.Context) error {

	var err resterrors.RestErr

	whereFrom := ctx.Param("where_from")
	whereTo := ctx.Param("where_to")

	if whereFrom == "" || whereTo == "" {
		err = resterrors.NewBadRequestError("Invalid parameters request")
		return ctx.JSON(err.StatusCode(), err)
	}

	bestRoute, err := c.travelService.GetBestRoute(strings.ToUpper(whereFrom), strings.ToUpper(whereTo))
	if err != nil {
		return ctx.JSON(err.StatusCode(), err)
	}

	response := viewmodel.TravelResponse{}
	mapperErr := c.mapper.From(bestRoute).To(&response)
	if mapperErr != nil {
		logger.Error("Error to mapper bestRouter to TravelResponse: ", mapperErr)
		err = resterrors.NewInternalServerError("Mapper response errror")
		return ctx.JSON(err.StatusCode(), err)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c Controller) handleAddNewRoute(ctx echo.Context) error {
	var err resterrors.RestErr

	var input viewmodel.TravelRequest
	parseErr := ctx.Bind(&input)
	if parseErr != nil {
		err = resterrors.NewBadRequestError("Invalid body request")
		return ctx.JSON(err.StatusCode(), err)
	}

	if input.WhereFrom == "" || input.WhereTo == "" || input.Price == 0 || input.WhereFrom == input.WhereTo {
		err = resterrors.NewBadRequestError("Invalid body request")
		return ctx.JSON(err.StatusCode(), err)
	}

	input.WhereFrom = strings.ToUpper(input.WhereFrom)
	input.WhereTo = strings.ToUpper(input.WhereTo)

	route := entity.Route{}
	mapperErr := c.mapper.From(input).To(&route)
	if mapperErr != nil {
		logger.Error("Error to mapper input request to route entity: ", mapperErr)
		err = resterrors.NewInternalServerError("Mapper response errror")
		return ctx.JSON(err.StatusCode(), err)
	}

	err = c.travelService.AddNewRoute(route)
	if err != nil {
		return ctx.JSON(err.StatusCode(), err)
	}

	return ctx.NoContent(http.StatusNoContent)
}
