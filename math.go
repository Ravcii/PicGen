package picgen

import (
	"golang.org/x/exp/constraints"
)

func Diff[T constraints.Float | constraints.Integer](from, to T) int {
	if from > to {
		return -int(from - to)
	}

	return int(to - from)
}
