package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("Hello, my name is", p.name)
}

type human interface {
	speak()
}

func main() {
	paul := person{"Paul McCartney"}

	// Works
	saySomething(&paul)

	// Doesn't works
	// saySomething(paul)
}

func saySomething(h human) {
	h.speak()
}
