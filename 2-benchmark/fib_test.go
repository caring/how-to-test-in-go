package fib_test

import (
	"fmt"
	"testing"

	fib "github.com/caring/test/2-benchmark"
)

// ExampleRecursive runs a slow recursive version of fib.
func ExampleRecursive() {
	fmt.Println(fib.Recursive(47))

	// Output: 2971215073
}

// ExampleLoop runs a medium iterative version of fib.
func ExampleLoop() {
	fmt.Println(fib.Loop(47))

	// Output: 2971215073
}

// ExampleFast runs a fast recursive version of fib.
func ExampleFast() {
	fmt.Println(fib.Fast(47))

	// Output: 2971215073
}

// BenchmarkFib will run all three fib implementations and print out the result.
func BenchmarkFib(b *testing.B) {
	benches := []struct {
		Name string
		Fn   func(uint) uint
		N    uint
	}{
		{"Recursive", fib.Recursive, 10},
		{"Loop", fib.Loop, 10},
		{"Fast", fib.Fast, 10},
	}

	for _, tt := range benches {
		b.Run(tt.Name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				tt.Fn(tt.N)
			}
		})
	}
}
