package circleinmem

import (
	domcircle "go-ddd/domain/circle"
	domuser "go-ddd/domain/user"
)

type Builder struct {
	m CircleModel
}

func (b *Builder) ID(id domcircle.CircleID)       { b.m.ID = id.String() }
func (b *Builder) Name(n domcircle.CircleName)    { b.m.Name = n.String() }
func (b *Builder) Owner(u *domuser.User) {
	if u != nil {
		b.m.OwnerID = u.ID().String()
	}
}
func (b *Builder) Members(ms []*domuser.User) {
	ids := make([]string, 0, len(ms))
	for _, u := range ms {
		if u != nil {
			ids = append(ids, u.ID().String())
		}
	}
	b.m.MemberIDs = ids
}


func (b * Builder) Build() CircleModel { return b.m }