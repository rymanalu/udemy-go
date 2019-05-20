package main

import "fmt"

func main() {
	a := make(chan int, 2)
	a <- 42
	a <- 43
	fmt.Println(<-a)
	fmt.Println(<-a)
	fmt.Printf("%T\n", a)

	// Send only type chan<- int
	b := make(chan<- int, 2)
	fmt.Printf("%T\n", b)
	// fmt.Println(<-b)

	// Receive only type chan<- int
	c := make(<-chan int, 2)
	fmt.Printf("%T\n", c)
}
