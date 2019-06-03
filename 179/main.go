package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type sqrtError struct {
	lat  float64
	long float64
	err  error
}

func (err sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", err.lat, err.long, err.err)
}

func main() {
	_, err := sqrt(-77)

	if err != nil {
		log.Println(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, sqrtError{42, 78, errors.New("error bro")}
	}

	return math.Sqrt(n), nil
}
