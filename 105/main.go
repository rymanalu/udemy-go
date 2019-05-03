package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.firstName, s.lastName)
}

func main() {
	sa1 := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			firstName: "Miss",
			lastName:  "Moneypenny",
		},
		ltk: true,
	}

	sa1.speak()

	sa2.speak()
}
