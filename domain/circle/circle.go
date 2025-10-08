package circle

import (
	user "go-ddd/domain/user"
)

const MaxMembers = 30

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
func (c *Circle) getNameForComparison() CircleName {
	return c.name
}

func (c *Circle) Join(member *user.User) error {
	if member == nil {
		return ErrCircleMemberEmpty
	}

	if c.IsFull() {
		return ErrCircleMembersTooMany
	}

	c.members = append(c.members, member)

	return nil
}

func (c *Circle) IsFull() bool {
	ownerLength := 1
	return ownerLength + len(c.members) >= MaxMembers
}

func (c *Circle) Notify(note CircleNotification) {
	membersCopy := append([]*user.User(nil), c.members...)

	note.ID(c.id)
	note.Name(c.name)
	note.Owner(c.owner)
	note.Members(membersCopy)
}