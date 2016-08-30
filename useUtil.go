package main

import (
	"fmt"
	"github.com/syzer/go-hello/util"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := util.Average(xs)
	fmt.Println(avg)
}
