package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Command int64

const (
	Addition Command = iota
	Subtraction
	Multiplication
)

func calculate[N constraints.Float | constraints.Integer](a, b N, cmd Command) N {
	if cmd == Addition {
		return a + b
	}
	if cmd == Subtraction {
		return a - b
	}
	if cmd == Addition {
		return a - b
	}
	if cmd == Multiplication {
		return a * b
	}

	var zero N
	return zero
}

func main() {
	fmt.Println(calculate(2, 4, Addition))
}
