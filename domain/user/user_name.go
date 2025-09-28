package user

import (
	"errors"
	"unicode/utf8"
)

type UserName string

const (
	MIN_LEN = 3
	MAX_LEN = 20
)

var (
	ErrNameEmpty = errors.New("user name is nil")
	ErrNameTooShort = errors.New("user name must be at least 3 characters")
	ErrNameTooLong = errors.New("user name must be at most 20 characters")
)

func NewUserName (value string) (UserName, error) {
	if value == "" {
		return "", ErrNameEmpty
	}

	length := utf8.RuneCountInString(value)

	if length < MIN_LEN {
		return "", ErrNameTooShort
	}

	if length > MAX_LEN {
		return "", ErrNameTooLong
	}

	name := UserName(value)
	return name, nil
}

func (un UserName) String() string {
	return string(un)
}

func (un UserName) Equals(other UserName) bool {
	return un == other
}