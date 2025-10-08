package circle

import (
	"strings"
)

type CircleID string
func (id CircleID) String() string {
	return string(id)
}

func NewCircleID(value string) (CircleID, error){
	v := strings.TrimSpace(value)
	if v == "" {
		return "", ErrCircleIDEmpty
	}

	return CircleID(v), nil
}