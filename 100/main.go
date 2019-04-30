package main

import "fmt"

func main() {
	u := struct {
		name string
		age  int
	}{
		name: "John",
		age:  40,
	}

	fmt.Println(u)
	fmt.Println(u.name, u.age)
}
