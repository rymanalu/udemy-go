package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2}
	fmt.Println("unsorted xi:", xi)
	sort.Ints(xi)
	fmt.Println("sorted xi:", xi)

	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry"}
	fmt.Println("unsorted xs:", xs)
	sort.Strings(xs)
	fmt.Println("sorted xs:", xs)
}
