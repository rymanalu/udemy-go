package main

import "fmt"

func main() {
	fib := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 35, 56}

	fmt.Println("fib sum all:", sum(fib...))

	fmt.Println("fib sum even:", sumEven(sum, fib...))

	fmt.Println("fib sum odd:", sumOdd(sum, fib...))
}

func sum(n ...int) int {
	total := 0

	for _, v := range n {
		total += v
	}

	return total
}

func sumEven(f func(n ...int) int, n ...int) int {
	var xn []int

	for _, v := range n {
		if v%2 == 0 {
			xn = append(xn, v)
		}
	}

	return f(xn...)
}

func sumOdd(f func(n ...int) int, n ...int) int {
	var xn []int

	for _, v := range n {
		if v%2 != 0 {
			xn = append(xn, v)
		}
	}

	return f(xn...)
}
