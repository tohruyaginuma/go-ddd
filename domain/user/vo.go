package user

import (
	"errors"
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
)

var (
	ErrEmptyUserName = errors.New("username cannot be empty")
	ErrTooShortUserName = errors.New("user name must be at least 3 characters")
)

type UserID uuid.UUID
type UserName string

func (id UserID) String() string {
	return uuid.UUID(id).String()
}

func (name UserName) String() string {
	return string(name)
}

func NewUserID() UserID{
    return UserID(uuid.New())
}

func ParseUserID(s string) (UserID, error) {
    u, err := uuid.Parse(s)
    if err != nil {
        return UserID{}, err
    }

    return UserID(u), nil
}

func NewUserName(s string) (UserName, error){
    s = strings.TrimSpace(s)

    if s == "" {
        return UserName(""), ErrEmptyUserName
    }

    if utf8.RuneCountInString(s) < 3 {
        return UserName(""), ErrTooShortUserName
    }

    return UserName(s), nil
}