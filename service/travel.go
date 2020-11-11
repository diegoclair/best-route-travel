package service

import (
	"github.com/diegoclair/best-route-travel/domain/contract"
	"github.com/diegoclair/best-route-travel/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type travelService struct {
	svc *Service
}

//newTravelService return a new instance of the service
func newTravelService(svc *Service) contract.TravelService {
	return &travelService{
		svc: svc,
	}
}

func (s *travelService) GetBestRoute(whereFrom, whereTo string) (bestRoute entity.BestRoute, err resterrors.RestErr) {

	return bestRoute, nil
}
