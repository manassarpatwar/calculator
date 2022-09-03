package main

import "fmt"

type Command int64

const (
	Addition Command = iota
	Subtraction
)

func main() {
	fmt.Println("Hello World!")
}
