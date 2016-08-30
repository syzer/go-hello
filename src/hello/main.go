package main

import (
	"fmt"
	"github.com/syzer/go-hello/src/stata"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := stata.Average(xs)
	fmt.Println(avg)
}
