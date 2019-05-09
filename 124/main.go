package main

import "fmt"

func main() {
	x := 69
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", &x)

	y := &x
	fmt.Println(y)
	fmt.Println(*y)
	fmt.Println(*&x)

	*y = 70
	fmt.Println(x)
}
