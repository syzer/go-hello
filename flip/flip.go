package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Print("test", text)
	fmt.Println(flip(text))
}

func flip(s string) string {
	switch s {
	case "┬─┬ノ(º_ºノ)\n":
		return "(╯°□°）╯︵ ┻━┻"
	default:
		return "┬─┬ノ(º_ºノ)"
	}
}
