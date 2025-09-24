package user

import (
	"fmt"
)

type userService struct {}

func NewUserService() userService {
    return userService{}
}

func (us userService) Exists(un UserName) bool {
    // TODO: Implement
    return false
}

func CreateUser(userName string)  {
    un, err := NewUserName(userName)

    if err != nil {
        panic(err)
    }

    var userService = NewUserService()

    if userService.Exists(un) {
        panic(fmt.Sprintf("%s already exists.", userName))
    }

    // TODO: Create
}