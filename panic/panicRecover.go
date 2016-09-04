package main

import (
	"fmt"
)

// 0
// <nil>
// 3
func main() {
	fmt.Println(divide(1, 0))
	fmt.Println(divide(6, 2))

	withPanic()

	x := 0
	changeValue(&x)
	fmt.Println("x is:", x)
}

func divide(a, b int) int {
	defer func() {
		fmt.Println(recover())
	}()
	return a / b
}

// PANIC
func withPanic() {
	defer func() {
		fmt.Println(recover())
	}()
	panic("PANIC")
}

func changeValue(x *int) {
	*x = 666
}
