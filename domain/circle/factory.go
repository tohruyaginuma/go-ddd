package circle

import (
	user "go-ddd/domain/user"
)

type Factory interface {
	Create(name CircleName, owner *user.User) (*Circle, error)
}