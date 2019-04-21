package main

import "fmt"

func main() {
	x := []int{0, 1, 1, 2, 3, 5}

	for i, v := range x {
		fmt.Println(i, v)
	}
}
