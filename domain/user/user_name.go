package user

import (
	"unicode/utf8"
)

type UserName string

const (
	MIN_LEN = 3
	MAX_LEN = 20
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