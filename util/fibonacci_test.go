package util_test

import (
	"testing"
	"fmt"
	"github.com/syzer/go-hello/util"
)

func TestFibonacci(t *testing.T) {
	t.SkipNow()
	test_case := util.Fibonacci()
	fmt.Println(test_case)
}
