package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(flip(text))
}

// "(╯°□°）╯︵ ┻━┻"
// "┬─┬ノ(º_ºノ)"
func flip(s string) string {
	if strings.Contains(s, "┬─┬") {
		s = strings.Replace(s, "┬─┬", "┻━┻", -1)
		s = strings.Replace(s, "ノ(º_ºノ)", "(╯°□°）╯", -1)
	} else {
		s = strings.Replace(s, "┻━┻", "┬─┬", -1)
		s = strings.Replace(s, "(╯°□°）╯", "ノ(º_ºノ)", -1)
	}
	return s
}
