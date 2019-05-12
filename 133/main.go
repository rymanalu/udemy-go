package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	fmt.Println("unsorted", xi)
	sort.Ints(xi)
	fmt.Println("sorted", xi)

	fmt.Println("----------------------------------")

	xs := []string{"John", "Paul", "George", "Ringo"}
	fmt.Println("unsorted", xs)
	sort.Strings(xs)
	fmt.Println("sorted", xs)
}
