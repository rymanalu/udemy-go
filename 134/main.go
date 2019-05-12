package main

import (
	"fmt"
	"sort"
)

type person struct {
	firstName string
	age       int
}

type byAge []person

func (p byAge) Len() int {
	return len(p)
}
func (p byAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p byAge) Less(i, j int) bool {
	return p[i].age < p[j].age
}

type byName []person

func (p byName) Len() int {
	return len(p)
}
func (p byName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p byName) Less(i, j int) bool {
	return p[i].firstName < p[j].firstName
}

func main() {
	p1 := person{"John", 40}
	p2 := person{"Paul", 77}
	p3 := person{"George", 58}
	p4 := person{"Ringo", 79}

	people := []person{p1, p2, p3, p4}

	fmt.Println("unsorted", people)
	sort.Sort(byAge(people))
	fmt.Println("sorted by age", people)
	sort.Sort(byName(people))
	fmt.Println("sorted by name", people)
}
