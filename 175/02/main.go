package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var errNorgateMath = errors.New("Norgate math: square root of negative number")

func main() {
	fmt.Printf("%T\n", errNorgateMath)

	_, err := sqrt(-69)

	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errNorgateMath
	}

	return math.Sqrt(f), nil
}
