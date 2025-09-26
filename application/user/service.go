package user

import (
	"fmt"
	domain "go-ddd/domain/user"
)

type Service struct {
	repo domain.Repository
	domainSrv domain.Service
}

func New(repo domain.Repository, domainSrv domain.Service) *Service {
	return &Service{
		repo: repo,
		domainSrv: domainSrv,
	}
}

func (s *Service) Register(name domain.UserName) error {
	user, err := domain.NewUser(name)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	isExist, err:= s.domainSrv.Exists(user.Name()); 

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if isExist {
		s.repo.Save(user)
	}

	return nil
}