package circle

import user "go-ddd/domain/user"

type CircleNotification interface {
	ID(id CircleID)
	Name(n CircleName)
	Owner(u *user.User)
	Members(ms []*user.User)
}