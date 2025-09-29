package user

type UserID string


func NewUserID(value string) (UserID, error) {
	if value == "" {
		return "", ErrUserIDEmpty
	}

	return UserID(value), nil
}

func (id UserID) String() string {
	return string(id)
}

func (id UserID) Equals(other UserID) bool {
	return id == other
}

