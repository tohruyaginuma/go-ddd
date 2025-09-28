package user

import domain "go-ddd/domain/user"

type UserDTO struct {
	ID string
	Name string
}

func FromDomain(u domain.User) UserDTO {
	return UserDTO{
		ID:	u.ID().String(),
		Name: u.Name().String(),
	}
}

func FromDomainSlice(us []domain.User) []UserDTO {
	users := make([]UserDTO, 0, len(us))
	
	for _, u := range us {
		users = append(users, FromDomain(u))
	}

	return users
}