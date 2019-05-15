package main

import "fmt"

func doSomething(n int) int {
	return n * 5
}

func main() {
	ch := make(chan int)

	go func() {
		ch <- doSomething(5)
	}()

	fmt.Println(<-ch)
}
