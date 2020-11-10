package service

import (
	"github.com/diegoclair/best-route-travel/domain/contract"
	"github.com/diegoclair/best-route-travel/domain/entity"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type userService struct {
	svc *Service
}

//newUserService return a new instance of the service
func newUserService(svc *Service) contract.UserService {
	return &userService{
		svc: svc,
	}
}

func (s *userService) SignIn(user entity.User) (entity.User, resterrors.RestErr) {

	// user, err := s.svc.db.User().FindByEmailAndPassword(user.Email, user.Password)
	// if err != nil {
	// 	return entity.User{}, err
	// }

	return user, nil
}
