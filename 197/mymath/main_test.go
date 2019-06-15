package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	avg := CenteredAvg([]int{1, 4, 6, 8, 100})

	if avg != 6 {
		t.Errorf("Expected %v, got %v", 6, avg)
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{0, 8, 10, 1000})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{9000, 4, 10, 8, 6, 12}))
	// Output:
	// 9
}
