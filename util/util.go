package util

// Math average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}


type Message struct {
	str  string
	wait chan bool
}

func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()                        //copy input1 to channel
	go func() {
		for {
			c <- <-input2
		}
	}()                        //copy input2 to channel
	return c
}

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		sum := a + b
		a = b
		b = sum
		return sum
	}
}
