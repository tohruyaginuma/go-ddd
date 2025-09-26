package user

type UserUpdateCommand struct {
	ID	string
	Name	*string
	MailAddress *string
}

type UserDeleteCommand struct {
	ID	string
}