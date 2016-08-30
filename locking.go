package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	// TODO
	c := fanIn(boring("John"), boring("Ann"))
	for i := 0; i < 20; i++ {
		msg1 := <-c; fmt.Println(msg1)
		//msg2 := <-c; fmt.Println(msg2.str)
		//msg1.wait <- true
		//msg2.wait <- true
		waitForIt := make(chan bool)
		//time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond) <- waitForIt
		c <- Message{fmt.Sprintf("%s %d", msg1, i), waitForIt }

	}
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()                        //copy input1 to channel
	go func() {
		for {
			c <- <-input2
		}
	}()                        //copy input2 to channel
	return c
}

// function that returns a chennel
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
