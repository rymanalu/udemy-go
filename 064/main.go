package main

import (
	"fmt"
	"time"
)

func main() {
	yearOfBirthDate := 1995

	for {
		if yearOfBirthDate > time.Now().Year() {
			break
		}

		fmt.Println(yearOfBirthDate)

		yearOfBirthDate++
	}
}
