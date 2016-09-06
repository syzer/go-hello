package main

import (
	"fmt"
	"io/ioutil"
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
