package inmemory

import (
	"context"
	domain "go-ddd/domain/user"
	"sync"
)

type Repository struct {
	mu	sync.RWMutex
	store map[domain.UserID] domain.User
}

func NewRepository() *Repository {
	return &Repository{
		store: make(map[domain.UserID]domain.User),
	}
}

func (r *Repository) FindByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if target, ok := r.store[id]; ok {
		return domain.NewUserWithID(target.ID(), target.Name()), nil
	}

	return nil, domain.ErrUserNotFound
}

func (r *Repository) FindByName(ctx context.Context, name domain.UserName) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, v := range r.store {
		if v.Name() == name {
			return domain.NewUserWithID(v.ID(), v.Name()), nil
		}
	}

	return nil, domain.ErrUserNotFound
}


func (r *Repository) FindAll(ctx context.Context) ([]domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]domain.User, 0, len(r.store))

	for _, v := range r.store {
		u := domain.NewUserWithID(v.ID(), v.Name())

		users = append(users, *u)
	}

	return users, nil
}
func (r *Repository) Save(ctx context.Context, user *domain.User) (error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store[user.ID()] = *domain.NewUserWithID(user.ID(), user.Name())

	return nil
}

func (r *Repository) Delete(ctx context.Context, user *domain.User) (error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if _, ok := r.store[user.ID()]; !ok {
		return domain.ErrUserNotFound	
	}

	
	delete(r.store, user.ID())
	return nil	
}