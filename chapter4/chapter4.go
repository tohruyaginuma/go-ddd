package chapter4

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"
)

type UserId string
type UserName string

func (id UserId) String() string {
    return string(id)
}

func (name UserName) String() string {
    return string(name)
}

func NewUserId () UserId{
    return UserId(uuid.NewString())
}

func NewUserName (s string) (UserName, error){
    s = strings.TrimSpace(s)

    if s == "" {
        return "", errors.New("username cannot be empty")
    }

    if utf8.RuneCountInString(s) < 3 {
        return "", errors.New("user name must be at least 3 characters")
    }

    return UserName(s), nil
}

type User struct {
	id	UserId
	name UserName
}

func NewUser(name UserName) (User, error){
    return User{
        id: NewUserId(),
        name: name,
    }, nil
}

func (u User) ID() UserId {
    return u.id
}

func (u User) Name() UserName {
    return u.name
}

func Execute() {
    un, err := NewUserName("toru")
    if err != nil {
        panic(err)
    }
    toru, err := NewUser(un)
    if err != nil {
        panic(err)
    }

    fmt.Println(toru.ID())
    fmt.Println(toru.Name())
    fmt.Println(toru.id)
    fmt.Println(toru.name)
}