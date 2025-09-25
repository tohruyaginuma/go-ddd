package usersvc

import (
	"fmt"
	domain "go-ddd/domain/user"
)

type Service struct{
	repo domain.Repository
	userSvc domain.Service
}

func New(repo domain.Repository, userSvc domain.Service) Service {
	return Service{
		repo: repo,
		userSvc: userSvc,
	} 
}

func (s *Service) CreateUser(userName string) (*domain.User, error) {
    un, err := domain.NewUserName(userName)

    if err != nil {
        return nil, err
    }

    if s.userSvc.Exists(un) {
		return nil, fmt.Errorf("%s already exists", userName)
    }
	name, err := domain.NewUserName(userName)

	if err != nil {
		return nil, err
	}
    user, err := domain.NewUser(name)

	if err != nil {
		return nil, err
	}
	return user, nil
}