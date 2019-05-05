package main

import "fmt"

func main() {
	fn := foo()
	fmt.Println(fn())

	fmt.Println(foo()())
}

func foo() func() int {
	return func() int {
		return 69
	}
}
