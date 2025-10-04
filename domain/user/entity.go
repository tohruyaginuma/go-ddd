package user

type User struct {
	id UserID
	name UserName
}

// ファクトリを用意することでこちらは不要になる。
// func NewUser(name UserName) (*User, error) {
// 	raw := uuid.NewString()

// 	id, err := NewUserID(raw)

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to generate user id: %w", err)
// 	}
	
// 	return &User{
// 		id: id,
// 		name: name,
// 	}, nil
// }

func NewUserWithID(id UserID, name UserName) *User {
	return &User{
		id: id,
		name: name,
	}
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Name() UserName {
	return u.name
}

func (u *User) ChangeName(name UserName) {
	u.name = name
}