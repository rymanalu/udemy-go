package word

import (
	"fmt"
	"testing"

	"github.com/rymanalu/udemy-go/196/quote"
)

func TestUseCount(t *testing.T) {
	c := UseCount(quote.SunAlso)

	if c["his"] != 20 {
		t.Errorf(`Expected number of "his" is %v, got %v`, 20, c["his"])
	}
}

func TestCount(t *testing.T) {
	c := Count("John Lennon")

	if c != 2 {
		t.Errorf(`Expected length of John Lennon is %v, got %v`, 2, c)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func ExampleUseCount() {
	fmt.Println(UseCount("The Beatles, Bob Dylan, The Rolling Stones"))
	// Output:
	// map[Beatles,:1 Bob:1 Dylan,:1 Rolling:1 Stones:1 The:2]
}

func ExampleCount() {
	fmt.Println(Count("John Lennon"))
	// Output:
	// 2
}
