package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	c := generator4("dupa")
	timeout := time.After(3 * time.Second)

	for {
		select {
		case s := <-c:
			log.Println(s)
		case <-timeout:
			log.Println("Done")
			return
		}
	}
}

// func that returns channel
// returned channel is receive only
func generator4(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("calling %s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		}
	}()
	return c
}
