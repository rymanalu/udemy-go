package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type musician struct {
	person
	alive bool
}

func main() {
	m1 := musician{
		person: person{
			firstName: "John",
			lastName:  "Lennon",
			age:       40,
		},
		alive: false,
	}

	m2 := musician{
		person: person{
			firstName: "Paul",
			lastName:  "McCartney",
			age:       38,
		},
		alive: true,
	}

	fmt.Println(m1)
	fmt.Println(m2)

	fmt.Println(m1.person.firstName, m1.lastName, m1.age, m1.alive)
	fmt.Println(m2.person.firstName, m2.person.lastName, m2.person.age, m2.alive)
}
