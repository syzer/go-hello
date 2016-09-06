package util_test

import (
	"fmt"
	"github.com/syzer/go-hello/util"
	"testing"
)

func TestFibonacci(t *testing.T) {
	t.SkipNow()
	test_case := util.Fibonacci()
	fmt.Println(test_case)
}
