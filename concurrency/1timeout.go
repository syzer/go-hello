package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {

	callMeMaybe("dupa")

}

func callMeMaybe(msg string) {
	for i := 0; i < 5; i++ {
		log.Println("calling ", msg, i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}
