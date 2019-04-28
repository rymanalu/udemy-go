package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{
		firstName: "John",
		lastName:  "Lennon",
		age:       40,
	}

	p2 := person{
		firstName: "Paul",
		lastName:  "McCartney",
		age:       38,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstName, p1.lastName, p1.age)
	fmt.Println(p2.firstName, p2.lastName, p2.age)
}
