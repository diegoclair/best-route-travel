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
	UserService(svc *Service) contract.UserService
	UploadService(svc *Service) contract.UploadService
}

type serviceManager struct {
	svc *Service
}

// NewServiceManager return a service manager instance
func NewServiceManager() Manager {
	return &serviceManager{}
}

func (s *serviceManager) UserService(svc *Service) contract.UserService {
	return newUserService(svc)
}

func (s *serviceManager) UploadService(svc *Service) contract.UploadService {
	return newUploadService(svc)
}
