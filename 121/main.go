package main

import "fmt"

func main() {
	foo(func() string {
		return "world"
	})
}

func foo(f func() string) {
	x := f()

	fmt.Println("Hello,", x)
}
