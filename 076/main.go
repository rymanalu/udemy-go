package main

import "fmt"

func main() {
	x := []int{0, 1, 1, 2, 3, 5}
	fmt.Println(x)
	x = append(x, 8, 13, 21, 35, 56)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
