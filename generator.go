package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := boring("boring")
	otherChann := boring("Other chann")

	for i := 0; i < 10; i++ {
		// HERE SYNHRONIZATION!!!!
		// no need for promises
		fmt.Printf("you say %q\n", <-c)
		fmt.Printf("you say %q\n", <-otherChann)
	}
	fmt.Println("DONE!")
}

// function that returns a channel
func boring(msg string) <-chan string {
	// => returns recive only channels of strings
	c := make(chan string)
	go func() {
		// goroutine
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	// returns a channel
	return c
}
