package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("called from assigned func")
	}

	f()
}
