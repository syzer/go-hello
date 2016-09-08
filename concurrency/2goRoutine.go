package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {

	go callMeMaybe2("dupa")
	// otherwise main would return
	time.Sleep(2 * time.Second)
}

func callMeMaybe2(msg string) {
	for i := 0; i < 5; i++ {
		log.Println("calling ", msg, i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}
