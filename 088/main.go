package main

import "fmt"

func main() {
	x := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Penny", "Helooooo, James"},
	}

	for i, slice := range x {
		fmt.Printf("record: %v\n", i)
		for j, v := range slice {
			fmt.Printf("\tindex position: %v\tvalue: %v\n", j, v)
		}
	}
}
