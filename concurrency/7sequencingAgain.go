package main

import (
	//"fmt"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Message struct {
	// normal message passing
	str string
	// flag when to stop blocking
	wait chan bool
}

//func main() {
//	c := make(chan string)
//
//	for i := 0; i < 5; i++ {
//		msg1 := <-c
//		log.Println(msg1.str)
//
//		msg2 := <-c
//		log.Println(msg2.str)
//
//		msg1.wait <- true
//		msg2.wait <- true
//	}
//
//	log.Println("Done")
//}
//
//func generator2(msg string) <-chan string {
//	c := make(chan string)
//	waiting := make(chan bool)
//	c <- Message{fmt.Sprintf("%s %d", msg, i), waiting}
//	time.Sleep(time.Duration(rand.Intn(800)) * time.Millisecond)
//	<-waiting
//
//	go func() {
//		for i := 0; i < 10; i++ {
//			c <- fmt.Sprintf("calling %s %d", msg, i)
//			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
//		}
//	}()
//	return c
//}
