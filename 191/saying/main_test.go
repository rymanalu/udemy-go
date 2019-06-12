package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Paul")

	expected := "Hello, Paul"

	if s != expected {
		t.Errorf(`Expected "%v" got "%v"`, expected, s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("George"))
	// Output:
	// Hello, George
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Ringo")
	}
}
