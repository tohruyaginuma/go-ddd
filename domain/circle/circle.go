package circle

import (
	user "go-ddd/domain/user"
)

type Circle struct {
	id CircleID
	name CircleName
	owner *user.User
	members []*user.User
}

func NewCircle(id CircleID, name CircleName, owner *user.User, members []*user.User) (*Circle, error) {
	if id == "" {
		return nil, ErrCircleIDEmpty
	}

	if name == "" {
		return nil, ErrCircleNameEmpty
	}

	if owner == nil {
		return nil, ErrCircleOwnerEmpty
	}

	if members == nil {
		members = make([]*user.User, 0)
	}

	return &Circle{
		id: id,
		name: name,
		owner: owner,
		members: members,
	}, nil
}

func (c *Circle) ID() CircleID { return c.id } 
func (c *Circle) Name() CircleName { return c.name } 
func (c *Circle) Owner() *user.User { return c.owner } 
func (c *Circle) Members() []*user.User { 
	out := make([]*user.User, len(c.members))
	copy(out, c.members)
	return out
} 