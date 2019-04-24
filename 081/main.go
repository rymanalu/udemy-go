package main

import "fmt"

func main() {
	m := map[string]int{
		"Roni":   24,
		"Kesbor": 69,
	}
	fmt.Println(m)

	delete(m, "Roni")
	fmt.Println(m)

	// No error thrown
	delete(m, "Random")
	fmt.Println(m)

	if v, ok := m["Kesbor"]; ok {
		fmt.Println(v)
		delete(m, "Kesbor")
	}
	fmt.Println(m)
}
