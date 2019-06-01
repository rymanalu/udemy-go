package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer fmt.Println("won't be printed")

	_, err := os.Open("no-file.txt")

	if err != nil {
		log.Fatalln("err happened", err)
	}
}
