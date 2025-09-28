package user

import "errors"

var (
	ErrNameEmpty = errors.New("user name is nil")
	ErrUserIDEmpty = errors.New("user id is empty")
	ErrNameTooShort = errors.New("user name must be at least 3 characters")
	ErrNameTooLong = errors.New("user name must be at most 20 characters")
	ErrUserNotFound = errors.New("user not found")
)