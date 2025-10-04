package inmemory

import (
	"context"
	"fmt"
	domain "go-ddd/domain/user"
	"sync/atomic"
)

type UserFactory struct {
	currentID int64
}

func NewUserFactory() *UserFactory{
	return &UserFactory{}
}

func (f *UserFactory) Create(ctx context.Context, name domain.UserName) (*domain.User, error) {
	idNum := atomic.AddInt64(&f.currentID, 1)
	userID, err := domain.NewUserID(fmt.Sprintf("%d", idNum))

	if err != nil {
		return nil, err
	}

	return domain.NewUserWithID(userID, name), nil
}