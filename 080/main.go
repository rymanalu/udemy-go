package main

import "fmt"

func main() {
	m := map[string]int{
		"Roni":   24,
		"Kesbor": 69,
	}
	fmt.Println(m)

	fmt.Println(m["Roni"])

	fmt.Println(m["Random"])

	v, ok := m["Random"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Roni"]; ok {
		fmt.Println("This is the if print", v)
	}

	m["AA"] = 69
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
