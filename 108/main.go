package main

import "fmt"

func main() {
	f1 := func() {
		fmt.Println("anonymous func ran")
	}
	f1()

	f2 := func(n int) {
		fmt.Println("I love", n)
	}
	f2(69)
}
