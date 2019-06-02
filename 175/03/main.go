package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	_, err := sqrt(-69)

	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("Norgate math: square root of negative number: %v", f)
	}

	return math.Sqrt(f), nil
}
