package main

import "fmt"

func main() {
	x := []int{0, 1, 1, 2, 3, 5}
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
