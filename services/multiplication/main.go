package multiplication

import (
	"golang.org/x/exp/constraints"
)

func calculate[N constraints.Float | constraints.Integer](a, b N) N {
	return a * b
}
