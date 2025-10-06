package circle

import (
	"strings"
	"unicode/utf8"
)

const (
	minNameLength = 3
	maxNameLength = 20
)

type CircleName string

func (name CircleName) String() string {
	return string(name)
}

func NewCircleName(value string) (CircleName, error){
	v := strings.TrimSpace(value)
	if v == "" {
		return "", ErrCircleNameEmpty
	}

	n := utf8.RuneCountInString(v)

	if n < minNameLength {
		return "", ErrCircleNameTooShort
	}

	if n > maxNameLength {
		return "", ErrCircleNameTooLong
	}

	return CircleName(v), nil
}