package pingroute

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds user handlers
type Controller struct {
}

//NewController to handle requests
func NewController() *Controller {
	once.Do(func() {
		instance = &Controller{}
	})
	return instance
}

// handlePing - handle a Ping request
func (s *Controller) handlePing(c echo.Context) error {
	return c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
