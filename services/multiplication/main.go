package multiplication

import (
	"golang.org/x/exp/constraints"
)

func calculate[N constraints.Float | constraints.Integer](a, b N) N {
	if a == 0 {
		return 0
	}
	if b == 0 {
		return 0
	}
	return a * b
}
