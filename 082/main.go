package main

import "fmt"

func main() {
	x := [5]int{100, 90, 80, 70, 60}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}
