package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer fmt.Println("still be printed")

	_, err := os.Open("no-file.txt")

	if err != nil {
		log.Panicln("err happened", err)
	}
}
