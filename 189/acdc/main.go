// Package acdc provides ACME inc math solutions.
package acdc

// Sum adds an unlimited number of values of type int.
func Sum(xi ...int) int {
	sum := 0

	for _, x := range xi {
		sum += x
	}

	return sum
}
