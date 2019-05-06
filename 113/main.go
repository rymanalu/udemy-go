package main

import "fmt"

func main() {
	f := foo()
	fmt.Println("foo =", f)

	b1, b2 := bar()
	fmt.Println(b1, "is a", b2)
}

func foo() int {
	return 69
}

func bar() (int, string) {
	return 666, "Devil's number"
}
