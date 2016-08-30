package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		sum := a + b
		a = b
		b = sum
		return sum
	}
}

func main() {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
