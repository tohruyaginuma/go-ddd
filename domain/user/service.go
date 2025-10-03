package user

import (
	"context"
	"errors"
	"fmt"
)

type Service struct {
	ur Repository
}

func New(ur Repository) Service {
	return Service{
		ur: ur,
	}
}

func (s *Service) Exists(ctx context.Context, user User) (bool, error) {
	_, err := s.ur.FindByName(ctx, user.Name())

	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return false, nil
		}

		return false, fmt.Errorf("failed to find user by name: %w", err)
	}

	return true, nil
}