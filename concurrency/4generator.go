package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {

	// its like yield from service in python / js
	c := generator("dupa")

	for i := 0; i < 5; i++ {
		log.Printf("channel emmited msg: %v", <-c)
	}

	log.Println("Done")
}

// func that returns channel
// returned channel is receive only
func generator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("calling %s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
	}()
	return c
}
