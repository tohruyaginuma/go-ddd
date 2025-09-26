package user

import (
	"context"
	"errors"
	"fmt"
	domain "go-ddd/domain/user"
	"strings"
)

var (
	ErrEmptyName = errors.New("name is required")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)

type Service struct {
	repo domain.Repository
	domainSrv *domain.Service
}

func New(repo domain.Repository, domainSrv *domain.Service) *Service {
	return &Service{
		repo: repo,
		domainSrv: domainSrv,
	}
}

func (s *Service) Register(ctx context.Context, name string) error {	
	userName, err := domain.NewUserName(name)
	
	if err != nil {
		return fmt.Errorf("failed to create user name: %w", err)
	}

	exists, err:= s.domainSrv.Exists(ctx, userName); 

	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}

	if exists {
		return ErrUserAlreadyExists
	}

	user, err := domain.NewUser(userName)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := s.repo.Save(ctx, user); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}

func (s *Service) Get(ctx context.Context, id string) (*UserData, error) {
	userId, err := domain.NewUserID(id)
	
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err)
	}

	user, err := s.repo.FindByID(ctx, userId)

	if err != nil {
		return nil, ErrUserNotFound
	}

	return FromDomain(user), nil
}

func (s *Service) Update (ctx context.Context, command UserUpdateCommand) error {
	userId, err := domain.NewUserID(command.ID)

	if err != nil {
	return fmt.Errorf("invalid user id: %w", err)
	}

	user, err := s.repo.FindByID(ctx, userId)

	if  err != nil {
		return fmt.Errorf("find user by id: %w", err)
	}

	if  user == nil{
		return ErrUserNotFound 
	}

	if command.Name != nil {
		newName, err := domain.NewUserName(strings.TrimSpace(*command.Name))
		
		if err != nil{
			return fmt.Errorf("failed to create user name: %w", err)
		}

		exists, err := s.domainSrv.Exists(ctx, newName)

		if err != nil {
			return fmt.Errorf("failed to check if user exists: %w", err)
		}
	
		if exists {
			return ErrUserAlreadyExists
		}

		user.ChangeName(newName)
	}

	if command.MailAddress != nil {
		// Implements
	}

	if err := s.repo.Save(ctx, user); err != nil {
		return fmt.Errorf("save user: %w", err)
	}
	
	return nil
}

func (s *Service) Delete(ctx context.Context, command UserDeleteCommand) error {
	userID, err := domain.NewUserID(command.ID)

	if err != nil {
		return fmt.Errorf("invalid user id: %w", err)
	}

	user, err := s.repo.FindByID(ctx, userID)

	if err != nil {
		return fmt.Errorf("find user by id: %w", err)
	}

	if user == nil {
		// If the user does not exist, treat as successful deletion
		return nil
	}

	if err := s.repo.Delete(ctx, user); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}


	return nil
}