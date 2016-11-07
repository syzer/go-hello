package main

import (
	"fmt"
)

// 0
// <nil>
// 3
func main() {
	x := timeToPanic(0,1)

	fmt.Println("x is:", x)
}

func timeToPanic(a,b int) (c int) {
	defer func() {
		recover()
		// implicitly return 42, nornally would be 0
		c = 42
	}()
    b = 0
	c = a/b
	return
}