package user

import "context"

type Repository interface {
	FindByID(ctx context.Context,id UserID) (*User, error)
	FindByName(ctx context.Context,name UserName) (*User, error)
	Save(ctx context.Context, u *User) error
	Delete(ctx context.Context, u *User) error
}