package main

import "testing"

func TestSumInt(t *testing.T) {
	sum := sumInt(2, 3)

	if sum != 5 {
		t.Error("Expected 5, got", sum)
	}
}
