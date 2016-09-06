package main

import (
	"fmt"
	"github.com/syzer/go-hello/boring"
	"github.com/syzer/go-hello/ranges"
	"github.com/syzer/go-hello/util"
	"time"
)

func main() {

	c := make(chan string)
	fmt.Printf("Hello world!")
	fmt.Println("Listening")
	// couroutine
	go boring.Boring("boring ", c)

	for i := 0; i < 10; i++ {
		// receive expression is just a value
		fmt.Printf("You say: %q\n", <-c)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("Timeout bitches")

	// multiplexer
	d := util.FanIn(boring.Boring_channel("John"), boring.Boring_channel("Ann"))
	for i := 0; i < 20; i++ {
		fmt.Printf("Multiplexer say %q\n", <-d)
	}
	fmt.Println("DONE!")

	// fibonacci
	ranges.Main()

	// use util
	xs := []float64{1, 2, 3, 4}
	avg := util.Average(xs)
	fmt.Println(avg)
}
