package user

type UserGetCommand struct {
	ID string
}

type UserUpdateCommand struct {
	ID string
	Name string
}

type UserRegisterCommand struct {
	Name string
}

type UserDeleteCommand struct {
	ID string
	Name string
}