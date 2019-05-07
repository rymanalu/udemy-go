package main

import "fmt"

func main() {
	fmt.Println("foo", foo(1, 2, 3, 4, 5))

	fmt.Println("bar", bar([]int{6, 7, 8, 9, 10}))
}

func foo(n ...int) int {
	total := 0

	for _, v := range n {
		total += v
	}

	return total
}

func bar(n []int) int {
	return foo(n...)
}
