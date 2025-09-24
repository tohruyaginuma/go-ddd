package user

type User struct {
	id	UserID
	name UserName
}

func NewUser(name UserName) (*User, error){
    return &User{
        id: NewUserID(),
        name: name,
    }, nil
}

func (u *User) ID() UserID {
    return u.id
}

func (u *User) Name() UserName {
    return u.name
}