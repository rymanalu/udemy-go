package main

import "fmt"

func main() {
	x := make(chan int)

	go func() {
		x <- 69
		close(x)
	}()

	v, ok := <-x
	fmt.Println(v, ok)

	v, ok = <-x
	fmt.Println(v, ok)
}
