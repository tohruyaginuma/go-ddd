package user

import "context"

type Service struct {
	ur Repository
}

func New(ur Repository) Service {
	return Service{
		ur: ur,
	}
}

func (s *Service) Exists(ctx context.Context, user User) (bool, error) {
	duplicatedUser, err := s.ur.FindByName(ctx, user.Name())

	if err != nil {
		return false, err
	}

	return duplicatedUser != nil, nil
}