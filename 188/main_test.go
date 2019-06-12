package main

import "testing"

func TestSumInt(t *testing.T) {
	type test struct {
		numbers []int
		answer  int
	}

	tests := []test{
		test{[]int{7, 5}, 12},
		test{[]int{2, 4}, 6},
		test{[]int{53, 42}, 95},
	}

	for _, testData := range tests {
		x := sumInt(testData.numbers...)

		if x != testData.answer {
			t.Error("Expected", testData.answer, "got", x)
		}
	}
}
