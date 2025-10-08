package circle

import (
	"context"
	"errors"
)

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {

	return &Service{
		repo: repo,
	}
}

func (s *Service) Exists(ctx context.Context, circle Circle) (bool, error) {
	name := circle.getNameForComparison()
	_, err := s.repo.FindByName(ctx, name)

	if err == nil {
		return true, err
	}

	if errors.Is(err, ErrCircleNotFound) {
		return false, nil
	}

	return false, err
}
