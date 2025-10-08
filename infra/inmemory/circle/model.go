package circleinmem

type CircleModel struct {
	ID        string
	Name      string
	OwnerID   string
	MemberIDs []string
}