package main

import "fmt"

func main() {
	x := 0
	fmt.Println("main", &x)
	fmt.Println("main", x)
	foo(&x)
	fmt.Println("main", &x)
	fmt.Println("main", x)
}

func foo(x *int) {
	fmt.Println("foo", *x)
	fmt.Println("foo", x)
	*x = 69
	fmt.Println("foo", *x)
	fmt.Println("foo", x)
}
