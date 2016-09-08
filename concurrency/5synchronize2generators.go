package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {

	// its like yield from service in python / js
	c := generator2("dupa")
	d := generator2("kupa")

	for i := 0; i < 5; i++ {
		// here sequencing: B always need to wait for A
		log.Printf("A channel emmited msg: %v", <-c)
		log.Printf("B channel emmited msg: %v", <-d)
	}

	log.Println("Done")
}

// func that returns channel
// returned channel is receive only
func generator2(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("calling %s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
	}()
	return c
}
