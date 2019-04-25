package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["ron"] = []string{"Yeah", "Kolok", "AI"}

	delete(m, "moneypenny_miss")

	for k1, v1 := range m {
		fmt.Println("Record for:", k1)

		for i, v2 := range v1 {
			fmt.Printf("\tindex: %v\tvalue: %v\n", i, v2)
		}
	}
}
