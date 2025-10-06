package circle

import (
	user "go-ddd/domain/user"
)

type CircleFactory interface {
	Create(name CircleName, owner *user.User) (*Circle, error)
}