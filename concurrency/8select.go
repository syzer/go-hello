package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	// like stream.join
	c := fanIn2(generator3("dupa"), generator3("kupa"))

	for i := 0; i < 10; i++ {
		log.Printf("channel emmited msg: %v", <-c)
	}

	log.Println("Done")
}

// like in multiplexer but with one subroutine
// AKA garded if in dikstra command
func fanIn2(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			// eval all channels.. blocks till one channel is ready
			// once is ready it can proceed
			// no default => nobody can proceed (AKA blocking)
			// order is kinda random (unles you nest one more)
			select {
			case s := <-c1:
				c <- s
			case s := <-c2:
				c <- s
			}
		}
	}()
	return c
}

func generator3(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("calling %s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
	}()
	return c
}
