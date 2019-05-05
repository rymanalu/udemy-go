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
	fmt.Println("I am", s.firstName, s.lastName, "- the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.firstName, p.lastName, "- the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).firstName)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).firstName)
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
		},
		ltk: true,
	}
	sa1.speak()

	sa2 := secretAgent{
		person: person{
			firstName: "Miss",
			lastName:  "Moneypenny",
		},
		ltk: true,
	}
	sa2.speak()

	p1 := person{
		firstName: "John",
		lastName:  "Lennon",
	}
	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)
}
