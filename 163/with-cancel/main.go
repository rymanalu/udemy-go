package main

import (
	"context"
	"fmt"
)

func main() {
	ctx1 := context.Background()

	ctx2, cancel := context.WithCancel(ctx1)

	fmt.Println("context:\t", ctx2)
	fmt.Println("context err:\t", ctx2.Err())
	fmt.Printf("context type:\t%T\n", ctx2)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)

	fmt.Println("----------------------------------------------")

	cancel()

	fmt.Println("context:\t", ctx2)
	fmt.Println("context err:\t", ctx2.Err())
	fmt.Printf("context type:\t%T\n", ctx2)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
}
