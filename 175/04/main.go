package main

import (
	"fmt"
	"log"
	"math"
)

type norgateMathError struct {
	lat  float64
	long float64
	err  error
}

func (err norgateMathError) Error() string {
	return fmt.Sprintf("A norgate math error occurred: %v %v %v", err.lat, err.long, err.err)
}

func main() {
	_, err := sqrt(-69)

	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		err := fmt.Errorf("Norgate math redux: square root of negative number: %v", f)
		return 0, norgateMathError{50.2289, 99.4656, err}
	}

	return math.Sqrt(f), nil
}
