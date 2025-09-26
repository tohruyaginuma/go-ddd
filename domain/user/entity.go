package user

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	id UserID
	name UserName
}

func (u *User) ChangeName(name UserName) {
	u.name = name
} 

func (u *User) Name() UserName {
	return u.name
} 

func (u *User) ID() UserID {
	return u.id
} 

func NewUser(name UserName) (*User, error) {
	id, err := NewUserID(uuid.New().String())

	if err != nil {
		return nil, fmt.Errorf("failed to generate user id: %w", err)
	}

	return &User{
		id: id,
		name: name,
	}, nil
}

func NewtUserWithID(id UserID, name UserName) (*User, error) {
	return &User{
		id: id,
		name: name,
	}, nil
}

