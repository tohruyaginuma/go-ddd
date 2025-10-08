package circle

type createCommand struct {
	userID string
	name string
}

type joinCommand struct {
	userID string
	circleID string
}