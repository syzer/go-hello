package boring

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	fmt.Printf("Hello world!")
	fmt.Println("Listening")
	// couroutine
	go Boring("boring ", c)

	for i := 0; i < 10; i++ {
		// receive expression is just a value
		fmt.Printf("You say: %q\n", <-c)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("Timeout bitches")
}

func Boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		//		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

// function that returns a channel
func Boring_channel(msg string) <-chan string {
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
