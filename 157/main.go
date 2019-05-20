package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)

	go receive(c)

	fmt.Println("exit")
}

func send(c chan<- int) {
	c <- 69
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
