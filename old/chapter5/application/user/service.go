package user

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
		return nil, fmt.Errorf("check exists: %w", err)
	}

	if isExist {
		return nil, fmt.Errorf("%s already exists", userName) 
	}

    user, err := domain.NewUser(un)

	if err != nil {
		return nil, err
	}

	if _, err := s.repo.Save(ctx, user); err != nil {
		return nil, fmt.Errorf("save user: %w", err)
	}

	return user, nil
}