package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// handle any panics that might occur
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}
	
	data := string(bs)

	words := strings.Fields(data)

	fmt.Println("Number of words in file:", len(words)) // expected count: 10
}