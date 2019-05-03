package main

import "fmt"

func main() {
	defer foo()
	bar()
	baz()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func baz() {
	fmt.Println("baz")
}
