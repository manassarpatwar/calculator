package subtraction

import (
	"golang.org/x/exp/constraints"
)

type Command int64

func calculate[N constraints.Float | constraints.Integer](a, b N) N {
	return a - b
}
