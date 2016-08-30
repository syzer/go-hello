package boring

import (
	"fmt"
	"time"
	"math/rand"
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

