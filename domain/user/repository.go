package user

type Repository interface {
	FindByID(id UserID) (*User, error)
	FindByName(name UserName) (*User, error)
	Save(u *User) error
	Delete(u *User) error
}