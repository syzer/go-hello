package main

import "fmt"

func main() {
	const n = 10000

	leftmost := make(chan int)
	left := leftmost
	right := leftmost

	// they all wait for first num to be send
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	// go gophers go!
	go func(c chan int) {
		c <- 1
	}(right)

	fmt.Println(<-leftmost)
}

func f(left, right chan int) {
	left <- 1 + <-right
}
