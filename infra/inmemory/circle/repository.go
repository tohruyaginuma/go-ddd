package circleinmem

import (
	"context"
	"sync"

	domcircle "go-ddd/domain/circle"
)

type Repository struct {
	mu	sync.RWMutex
	store	map[string]CircleModel
}

func NewRepository() *Repository {
	return &Repository{store: make(map[string]CircleModel)}
}

func (r *Repository) Save(ctx context.Context, c *domcircle.Circle) error {
	b := &Builder{}
	c.Notify(b)
	
	model := b.Build()
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[model.ID] = model
	return nil
}