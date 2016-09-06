package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println("=> " + path)
		getContent(path)
		return nil
	})
}

func getContent(fileName string) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Cannot open file %s", fileName)
		return
	}
	str := string(bs)
	fmt.Println(str)
}
