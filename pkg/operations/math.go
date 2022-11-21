package operations

import (
	"golang.org/x/exp/constraints"
)

func Diff[T constraints.Float | constraints.Integer](from, to T) T {
	if from > to {
		return from - to
	}

	return to - from
}
