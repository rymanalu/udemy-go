package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, "and I am", p.age, "year(s) old")
}

func main() {
	p := person{
		first: "Paul",
		last:  "McCartney",
		age:   76,
	}

	p.speak()
}
