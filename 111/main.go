package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x", x)
	x++

	{
		y := 1
		fmt.Println("y", y)
	}

	fmt.Println("x", x)

	// ------------------------------------------------------------------------ //

	a := incrementor()
	b := incrementor()

	fmt.Println("a", a())
	fmt.Println("a", a())
	fmt.Println("a", a())
	fmt.Println("a", a())
	fmt.Println("b", b())
	fmt.Println("b", b())
}

func incrementor() func() int {
	var x int

	return func() int {
		x++

		return x
	}
}
