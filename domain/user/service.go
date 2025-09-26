package user

import "fmt"

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Exists(name UserName) (bool, error) {
	duplicatedUser, err := s.repo.FindByName(name)

	if err != nil {
		return false, fmt.Errorf("failed to check if user exists: %w", err)
	}

	return duplicatedUser != nil, nil
}