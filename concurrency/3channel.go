package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)

	go callMeMaybe3("dupa", c)

	for i := 0; i < 5; i++ {
		log.Printf("channel emmited msg: %v", <-c)
	}
	log.Println("Done")
}

func callMeMaybe3(msg string, c chan string) {
	for i := 0; i < 5; i++ {
		// pass expression!!
		c <- fmt.Sprintf("calling %s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}
