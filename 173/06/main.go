package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("won't be printed")

	_, err := os.Open("no-file.txt")

	if err != nil {
		panic(err)
	}
}
