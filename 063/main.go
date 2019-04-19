package main

import (
	"fmt"
	"time"
)

func main() {
	yearOfBirthDate := 1995

	for yearOfBirthDate <= time.Now().Year() {
		fmt.Println(yearOfBirthDate)

		yearOfBirthDate++
	}
}
