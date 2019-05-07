package main

import "fmt"

func main() {
	foo()
	defer bar()
	fmt.Println("printed after foo")
	fmt.Println("printed before bar")
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
