package main

import "fmt"

type newInt int

var x newInt

func main() {
	fmt.Println("x = ", x)
	fmt.Printf("type of x = %T\n", x)
	x = 42
	fmt.Println("x = ", x)
}
