package main

import "fmt"

func main() {
	f := foo()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func foo() func() int {
	n := 0

	return func() int {
		n++

		return n
	}
}
