package main

import (
	"fmt"
	"time"
	"github.com/syzer/go-hello/boring"
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
}
