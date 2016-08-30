package main

import (
	"testing"
	"fmt"
)

func TestFibonacci(t *testing.T) {
	t.SkipNow()
	test_case := Fibonacci()
	fmt.Println(test_case)
}
