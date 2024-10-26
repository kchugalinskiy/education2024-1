package factorial

import (
	"fmt"
	"testing"
)

// ....
func ExampleFactorial() {
	fmt.Println(Factorial(5))
	// Output: 120
}

func TestFactorial(t *testing.T) {
	tst := []struct {
		n int
		f int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
	}
	for _, tt := range tst {
		if f := Factorial(tt.n); f != tt.f {
			t.Errorf("Factorial(%d) = %d, want %d", tt.n, f, tt.f)
		}
	}
}

func TestFactorial_Zero(t *testing.T) {
	fmt.Println(Factorial(0))
	// Output: 1
}

func TestFactorial_One(t *testing.T) {
	fmt.Println(Factorial(1))
	// Output: 1
}

func TestFactorial_Five(t *testing.T) {
	fmt.Println(Factorial(5))
	// Output: 120
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(15)
	}
}
