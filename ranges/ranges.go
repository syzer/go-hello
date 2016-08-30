package ranges

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Main() {
	// le normal
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow := make([]int, 10)

	// just a index
	for i := range pow {
		// shift
		pow[i] = 1 << uint(i) // == 2**i
	}

	// just a value
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
