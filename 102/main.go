package main

import "fmt"

func main() {
	x := sum(0, 1, 1, 2, 3, 5, 8)
	fmt.Println("The total is", x)
}

func sum(n ...int) int {
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	sum := 0
	for i, v := range n {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}

	return sum
}
