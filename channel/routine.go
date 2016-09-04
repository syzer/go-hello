package main

import (
	"time"
	"fmt"
	"strconv"
)

var carNum = 0
var carName = ""
const sleepTime = 6

func makeCarChassis(stringChan chan string) {

	fmt.Println("Making car chassis")

	carNum++

	// Convert int into a string
	carName = "car #" + strconv.Itoa(carNum)

	// Send the carName onto the channel for the next
	stringChan <- carName

	time.Sleep(time.Millisecond * sleepTime)
}

// channel allows to pass data to coroutine
func addEngine(stringChan chan string) {

	// Receive the value passed on the channel
	car := <-stringChan

	fmt.Println("Adding Engine and forward", car, "to Wheels factory")

	// Send the carName onto the channel for the next
	stringChan <- carName

	time.Sleep(time.Millisecond * sleepTime)
}

func addWheels(stringChan chan string) {

	// Receive the value passed on the channel
	car := <-stringChan

	fmt.Println("Adding Wheels to", car, "and ship it!")

	time.Sleep(time.Millisecond * sleepTime)
}

func main() {

	// Make creates a channel that can hold a string
	// int channel intChan := make(chan int)
	stringChan := make(chan string)

	// Cycle through and make 3 cars
	for i := 0; i < 3; i++ {

		go makeCarChassis(stringChan)
		go addEngine(stringChan)
		go addWheels(stringChan)

		time.Sleep(time.Millisecond * sleepTime * 100)
	}

}
