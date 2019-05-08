package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("Area:", s.area())
}

func main() {
	s := square{
		side: 6,
	}

	info(s)

	c := circle{
		radius: 9,
	}

	info(c)
}
