package user

import domain "go-ddd/domain/user"

type UserDto struct {
	ID string
	Name string
}

func FromDomain(u domain.User) UserDto {
	return UserDto{
		ID:	u.ID().String(),
		Name: u.Name().String(),
	}
}

func FromDomainSlice(us []domain.User) []UserDto {
	users := make([]UserDto, 0, len(us))
	
	for _, u := range us {
		users = append(users, FromDomain(u))
	}

	return users
}