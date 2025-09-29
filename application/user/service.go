package user

import (
	"context"
	"fmt"
	domain "go-ddd/domain/user"
)

type Service struct {
	r domain.Repository
	us domain.Service
}

func New(r domain.Repository, us domain.Service) *Service {
	return &Service{
		r: r,
		us: us,
	}
}

func (s *Service) Get(ctx context.Context, cmd UserGetCommand) (UserDTO, error){
	id, err := domain.NewUserID(cmd.ID)

	if err != nil {
		return UserDTO{}, fmt.Errorf("invalid user ID: %w", err)
	}

	user, err := s.r.FindByID(ctx, id)

	if err != nil {
		return UserDTO{}, fmt.Errorf("failed to get all users: %w", err)
	}

	if user == nil {
		return UserDTO{}, fmt.Errorf("user not found")
	}

	return FromDomain(*user), nil
}

func (s *Service) GetAll(ctx context.Context) ([]UserDTO, error) {
	users, err := s.r.FindAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	return FromDomainSlice(users), nil
}

func (s *Service) Update(ctx context.Context, cmd UserUpdateCommand) (UserDTO, error) {
	id, err := domain.NewUserID(cmd.ID)
	if err != nil {
		return UserDTO{}, fmt.Errorf("invalid user ID: %w", err)
	}

	user, err := s.r.FindByID(ctx, id)
	if err != nil {
		return UserDTO{}, fmt.Errorf("failed to find user: %w", err)
	}
	if user == nil {
		return UserDTO{}, domain.ErrUserNotFound
	}

	if cmd.Name != "" {
		name, err := domain.NewUserName(cmd.Name)
		if err != nil {
			return UserDTO{}, fmt.Errorf("invalid user name: %w", err)
		}
		user.ChangeName(name)

		exists, err := s.us.Exists(ctx, *user)
		if err != nil {
			return UserDTO{}, fmt.Errorf("failed to check user existence: %w", err)
		}
		if exists {
			return UserDTO{}, fmt.Errorf("user already exists")
		}
	}

	if err := s.r.Save(ctx, user); err != nil {
		return UserDTO{}, fmt.Errorf("failed to save user: %w", err)
	}

	return FromDomain(*user), nil
}

func (s *Service) Register(ctx context.Context, cmd UserRegisterCommand) error {
	name, err := domain.NewUserName(cmd.Name)
	if err != nil {
		return fmt.Errorf("invalid user name: %w", err)
	}

	user, err := domain.NewUser(name)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := s.r.Save(ctx, user); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, cmd UserDeleteCommand) error {
	id, err := domain.NewUserID(cmd.ID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	user, err := s.r.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}
	if user == nil {
		// If user not found, do nothing (as in the C# sample)
		return nil
	}

	if err := s.r.Delete(ctx, user); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}