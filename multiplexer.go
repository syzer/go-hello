package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := fanIn(boring("John"), boring("Ann"))
	for i := 0; i < 20; i++ {
		fmt.Printf("you say %q\n", <-c)
	}
	fmt.Println("DONE!")
}

// fan in takes values from 2 channels and combines to one channel
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			// copy input1 to channel
			c <- <-input1
		}
	}()
	go func() {
		for {
			// copy input2 to channel
			c <- <-input2
		}
	}()
	return c
}

// function that returns a channel
func boring(msg string) <-chan string {
	//returns recive only channels of strings
	c := make(chan string)
	go func() {
		// goroutine
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c // returns a channel
}

