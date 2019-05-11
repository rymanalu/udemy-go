package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
	Age       int
}

// type person struct {
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	Age       int    `json:"age"`
// }

func main() {
	p1 := person{
		FirstName: "John",
		LastName:  "Lennon",
		Age:       40,
	}
	fmt.Println(p1)

	p2 := person{
		FirstName: "Paul",
		LastName:  "McCartney",
		Age:       76,
	}
	fmt.Println(p2)

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println(string(bs))
}
