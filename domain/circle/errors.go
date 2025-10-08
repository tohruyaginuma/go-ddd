package circle

import "errors"

var (
	ErrCircleMemberEmpty    = errors.New("member is empty")
	ErrCircleExists         = errors.New("circle already exists")
	ErrCircleNotExists		= errors.New("circle not exists")
	ErrCircleOwnerNotFound  = errors.New("circle owner not found")
	ErrCircleNotFound       = errors.New("circle not found")
	ErrCircleIDEmpty        = errors.New("circle id is empty")
	ErrCircleNameEmpty      = errors.New("circle name is empty")
	ErrCircleNameTooShort   = errors.New("circle name is too short (minimum 3 characters)")
	ErrCircleNameTooLong    = errors.New("circle name is too long (maximum 20 characters)")
	ErrCircleOwnerEmpty     = errors.New("circle owner is empty")
	ErrCircleMembersEmpty   = errors.New("circle members are empty")
	ErrCircleMembersTooMany = errors.New("too many circle members")
	ErrCircleMemberNotFound = errors.New("circle member not found")
)
