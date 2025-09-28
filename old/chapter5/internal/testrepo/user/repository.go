package inmemory

import (
	"context"
	domain "go-ddd/domain/user"
	"sync"
)

type Repository struct {
	mu sync.RWMutex
	store map[domain.UserID] *domain.User
}

func New() *Repository {
	return &Repository{
		store: make(map[domain.UserID] *domain.User),
	}
}

func (r *Repository) Find(ctx context.Context, un domain.UserName) (*domain.User, error)  {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.store {
		if u.Name() == un {
			clone := *u
			return &clone, nil
		}
	}
	
	return nil, nil
}

func (r *Repository) Save(ctx context.Context, u *domain.User) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	clone := *u
	r.store[u.ID()] = &clone
	
	return &clone, nil
}

