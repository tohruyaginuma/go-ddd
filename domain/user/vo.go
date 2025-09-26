package user

import (
	"errors"
	"strings"
	"unicode/utf8"
)

const (
	userNameMinLen = 3
	userNameMaxLen = 20
)

var (
	ErrUserIDEmpty = errors.New("user id cannot be empty")
	ErrUserNameEmpty = errors.New("user name cannot be empty")
	ErrUserNameLeast = errors.New("user name must be at least 3 characters long")
	ErrUserNameMost = errors.New("user name must be at most 20 characters long")
)

type UserID string
type UserName string

func (id UserID) String() string {
	return string(id)
}

func (name UserName) String() string {
	return string(name)
}

func NewUserID (id string) (UserID, error) {
	id = strings.TrimSpace(id)
	if id == "" {
		return UserID(""), ErrUserIDEmpty
	}

	return UserID(id),  nil
}

func NewUserName (name string) (UserName, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return UserName(""), ErrUserNameEmpty
	}

	nameLen := utf8.RuneCountInString(name)

	if nameLen < userNameMinLen {
		return UserName(""), ErrUserNameLeast
	}

	if nameLen > userNameMaxLen {
		return UserName(""), ErrUserNameMost
	}
	
	return UserName(name), nil
}

