package main

import (
	"fmt"
	"log"
)

func main() {
	// can be also flag
	quit := make(chan string)
	c := generator5("dupa", quit)

	for i := 10; i >= 0; i-- {
		log.Println(<-c)
	}

	quit <- "Stop!"

	// could also listen here for <-quit
	fmt.Println("Done")
}

func generator5(msg string, quit <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("%s", msg):
			case <-quit:
				log.Println("Done")
				//quit <- "i did cleanup db connection"
				return
			}
		}
	}()

	return c
}
