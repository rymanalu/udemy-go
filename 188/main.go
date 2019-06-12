package main

func main() {
	//
}

func sumInt(xi ...int) int {
	sum := 0

	for _, x := range xi {
		sum += x
	}

	return sum
}
