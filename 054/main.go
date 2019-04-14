package main

import "fmt"

func main() {
	for i := 33; i < 123; i++ {
		fmt.Printf("%v in ASCII is %c\n", i, i)
	}
}
