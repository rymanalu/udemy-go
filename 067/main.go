package main

import "fmt"

func main() {
	familyName := "Manalu"

	if familyName == "Panggabean" {
		fmt.Println("if", familyName)
	} else if familyName == "Manalu" {
		fmt.Println("else if", familyName)
	} else {
		fmt.Println("else")
	}
}
