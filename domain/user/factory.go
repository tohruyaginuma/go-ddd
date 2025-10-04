package user

import "context"

type Factory interface {
	Create(ctx context.Context, name UserName) (*User, error)
}