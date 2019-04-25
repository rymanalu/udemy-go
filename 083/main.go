package main

import "fmt"

func main() {
	x := []int{6, 4, 7, 3, 1, 2, 0, 9, 8, 5}

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", x)
}
