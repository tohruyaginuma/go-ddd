package circle

import (
	"context"
	"fmt"
	domCircle "go-ddd/domain/circle"
	domUser "go-ddd/domain/user"
)

type TxManager interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) error
}

type Service struct {
	factory domCircle.Factory
	domService domCircle.Service
	repoCircle domCircle.Repository
	repoUser domUser.Repository
	tx TxManager
}

func NewService(
	factory domCircle.Factory, 
	domService domCircle.Service, 
	repoCircle domCircle.Repository, 
	repoUser domUser.Repository,
	rx TxManager,
	) *Service {
    return &Service{
		factory: factory,
		domService: domService,
		repoCircle: repoCircle,
		repoUser: repoUser,
		tx: rx,
	}
}

func (s *Service) Create(ctx context.Context, command createCommand) error {
    run := func(ctx context.Context) error {
		ownerId, err := domUser.NewUserID(command.userID)

		if err != nil {
			return fmt.Errorf("invalid owner id: %w", err)
		}

		owner, err := s.repoUser.FindByID(ctx, ownerId)
	
		if err != nil {
			return fmt.Errorf("find owner by id: %w", err)
		}

		if owner == nil {
			return domCircle.ErrCircleOwnerNotFound
		}
	
		name, err := domCircle.NewCircleName(command.name)

		if err != nil {
			return fmt.Errorf("invalid circle name: %w", err)
		}

		circle, err := s.factory.Create(name, owner)

		if err != nil {
			return fmt.Errorf("failed to create circle: %w", err)
		}
	
		exists, err := s.domService.Exists(ctx, *circle)

		if err != nil {
			return fmt.Errorf("failed to check if circle exists: %w", err)
		}

		if exists {
			return domCircle.ErrCircleExists
		}

		if err = s.repoCircle.Save(ctx, circle); err != nil {
			return err
		}
		
        return nil
	}

	if s.tx != nil {
		return s.tx.Do(ctx, run)	
	}

	return run(ctx)
}

func (s *Service) Join(ctx context.Context, command joinCommand) error {
	run := func(ctx context.Context) error {
		memberID, err := domUser.NewUserID(command.userID)

		if err != nil {
			return fmt.Errorf("invalid member id: %w", err)
		}

		member, err := s.repoUser.FindByID(ctx, memberID)
		
		if err != nil {
			return fmt.Errorf("failed to find member by id: %w", err)
		}

		if member == nil {
			return domCircle.ErrCircleMemberNotFound
		}

		id, err := domCircle.NewCircleID(command.circleID)
		
		if err != nil {
			return fmt.Errorf("invalid circle id: %w", err)
		}

		circle, err := s.repoCircle.FindByID(ctx, id)

		if err != nil {
			return domCircle.ErrCircleNotFound
		}

		if circle == nil {
			return domCircle.ErrCircleNotExists
		}
		
		if circle.IsFull() {
			return domCircle.ErrCircleMembersTooMany
		}

		circle.AddMember(member)
		
		if err := s.repoCircle.Save(ctx, circle); err != nil {
			return fmt.Errorf("failed to save circle: %w", err)
		}

		return nil
	}

	if s.tx != nil {
		return s.tx.Do(ctx, run)
	}

	return run(ctx)
}