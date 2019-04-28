package main

import "fmt"

func main() {
	p1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "George",
		lastName:  "Harrison",
		age:       37,
	}

	fmt.Println(p1)
}
