package usersvc

import (
	"context"
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

func (s *Service) CreateUser(ctx context.Context, userName string) (*domain.User, error) {
    un, err := domain.NewUserName(userName)

    if err != nil {
        return nil, err
    }

	isExist, err := s.userSvc.Exists(ctx, un)

	if err != nil {
		return nil, fmt.Errorf("error occared") 
	}

	if isExist {
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