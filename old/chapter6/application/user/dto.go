package user

import (
	domain "go-ddd/domain/user"
)

type UserData struct {
	ID string
	Name string
}

func FromDomain(u *domain.User) *UserData {
	if u == nil {
		return nil
	}

	return &UserData{
		ID: u.ID().String(),
		Name: u.Name().String(),
	}
}