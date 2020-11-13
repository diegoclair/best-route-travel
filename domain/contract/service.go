package contract

import (
	"github.com/diegoclair/best-route-travel/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

// TravelService holds a travel service operations
type TravelService interface {
	GetBestRoute(whereFrom, whereTo string) (bestRoute entity.BestRoute, err resterrors.RestErr)
	AddNewRoute(route entity.Route) (err resterrors.RestErr)
}

// CommandLineService holds a cli service operations
type CommandLineService interface {
	RunCLI()
	InputNewFile(fileName string)
}
