package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	// like stream.join
	c := fanIn(generator2("dupa"), generator2("kupa"))

	for i := 0; i < 10; i++ {
		log.Printf("channel emmited msg: %v", <-c)
	}

	log.Println("Done")
}

// copies both channels to 1...
// so c1 and c2 are not blocked by each other
// order can be random.
func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1
		}
	}()
	go func() {
		for {
			c <- <-c2
		}
	}()
	return c
}

func generator2(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("calling %s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
	}()
	return c
}
