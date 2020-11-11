package service

import (
	"github.com/diegoclair/best-route-travel/domain/contract"
)

// Service holds the domain service repositories
type Service struct {
}

// New returns a new domain Service instance
func New() *Service {
	svc := new(Service)
	return svc
}

//Manager defines the services aggregator interface
type Manager interface {
	TravelService(svc *Service) contract.TravelService
	UploadService(svc *Service) contract.UploadService
	CommandLineService(svc *Service) contract.CommandLineService
}

type serviceManager struct {
	svc *Service
}

// NewServiceManager return a service manager instance
func NewServiceManager() Manager {
	return &serviceManager{}
}

func (s *serviceManager) TravelService(svc *Service) contract.TravelService {
	return newTravelService(svc)
}

func (s *serviceManager) UploadService(svc *Service) contract.UploadService {
	return newUploadService(svc)
}

func (s *serviceManager) CommandLineService(svc *Service) contract.CommandLineService {
	return newCommandLineService(svc)
}
