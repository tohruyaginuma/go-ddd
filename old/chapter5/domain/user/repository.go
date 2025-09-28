package user

import (
	"context"
)

type Repository interface {
    Save(ctx context.Context, u *User) (*User, error)
    Find(ctx context.Context, un UserName) (*User, error)
}
