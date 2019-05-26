package main

import "fmt"

func main() {
	// Using func literal...
	a := make(chan int)

	go func() {
		a <- 42
	}()

	fmt.Println(<-a)

	fmt.Println("---------------")

	// Using buffered channel...
	b := make(chan int, 1)

	b <- 42

	fmt.Println(<-b)
}
