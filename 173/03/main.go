package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f1, err := os.Create("log.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f1.Close()

	log.SetOutput(f1)

	f2, err := os.Open("no-file.txt")

	if err != nil {
		log.Println("err happened", err)
		return
	}

	defer f2.Close()

	fmt.Println("check log.txt")
}
