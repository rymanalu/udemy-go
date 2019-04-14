package main

import "fmt"

func main() {
	x := 40

	if x == 40 {
		fmt.Println("x is 40")
	} else if x == 41 {
		fmt.Println("x is 41")
	} else {
		fmt.Println("x is not 40 or 41")
	}
}
