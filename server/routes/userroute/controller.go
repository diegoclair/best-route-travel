package userroute

import (
	"net/http"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/best-route-travel/domain/contract"
	"github.com/diegoclair/best-route-travel/domain/entity"
	"github.com/diegoclair/best-route-travel/server/viewmodel"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/labstack/echo"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller is a interface to interact with services
type Controller struct {
	userService contract.UserService
	mapper      mapper.Mapper
}

//NewController to handle requests
func NewController(userService contract.UserService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			userService: userService,
			mapper:      mapper,
		}
	})
	return instance
}

func (c Controller) handleSignInUser(context echo.Context) error {

	var err resterrors.RestErr

	var vm viewmodel.SignInUserRequest
	parseErr := context.Bind(&vm)
	if parseErr != nil {
		err = resterrors.NewBadRequestError("Error parse body into User Login Request")
		return context.JSON(err.StatusCode(), err)
	}

	user, err := c.userService.SignIn(entity.User{})
	if err != nil {
		return context.JSON(err.StatusCode(), err)
	}

	response := user

	return context.JSON(http.StatusOK, response)
}
