package main

import "fmt"

type newInt int

var x newInt
var y int

func main() {
	fmt.Println("x = ", x)
	fmt.Printf("type of x = %T\n", x)
	x = 42
	fmt.Println("x = ", x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("type of y = %T\n", y)
}
