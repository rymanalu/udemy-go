package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 420
	b = 69
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
