package addition

import (
	"golang.org/x/exp/constraints"
)

func calculate[N constraints.Float | constraints.Integer](a, b N) N {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return a + b
}
