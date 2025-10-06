package circle

import "errors"

var (
	ErrCircleIDEmpty = errors.New("circle id is empty")
	ErrCircleNameEmpty = errors.New("circle name is empty")
	ErrCircleNameTooShort = errors.New("circle name is too short")
	ErrCircleNameTooLong = errors.New("circle name is too long")
	ErrCircleOwnerEmpty   = errors.New("owner is empty")
	ErrCircleMembersEmpty = errors.New("members are empty")
)
