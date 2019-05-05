package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("anonymous func ran")
	}()

	func(n int) {
		fmt.Println("I love", n)
	}(69)
}

func foo() {
	fmt.Println("foo ran")
}
