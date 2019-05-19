package main

import "fmt"

func main() {
	// Doesn't run (#1)
	/*a := make(chan int)
	a <- 42
	fmt.Println(<-a)*/

	// Run (#1)
	b := make(chan int)
	go func() {
		b <- 42
	}()
	fmt.Println(<-b)

	// Run (#2)
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)

	// Doesn't run (#2)
	/*d := make(chan int, 1)
	d <- 42
	d <- 43
	fmt.Println(<-d)*/

	// Run (#3)
	e := make(chan int, 2)
	e <- 42
	e <- 43
	fmt.Println(<-e)
	fmt.Println(<-e)
}
