package main

import (
	"fmt"

	"github.com/rymanalu/udemy-go/186/dog"
)

func main() {
	humanYears := 24

	fmt.Printf("%v in dog years is: %v\n", humanYears, dog.Years(humanYears))
}
