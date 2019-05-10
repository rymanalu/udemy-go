package main

import "fmt"

type person struct {
	name    string
	address string
}

func changeMe(p *person) {
	p.address = "Jakarta"
	// (*p).address = "Jakarta"
}

func main() {
	p1 := person{
		name:    "John Lennon",
		address: "London",
	}

	fmt.Println("p1:", p1)

	changeMe(&p1)

	fmt.Println("p1:", p1)
}
