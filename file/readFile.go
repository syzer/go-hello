package main

import (
	"io/ioutil"
	"fmt"
	"log"
)

func main() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := string(bs)
	fmt.Println(str)
}
