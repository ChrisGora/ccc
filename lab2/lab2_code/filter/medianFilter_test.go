package main

import (
	"os"
	"testing"
)

// Benchmark applies the filter to the ship.png b.N times.
// The time taken is carefully measured by go.
// The b.N  repetition is needed because benchmark results are not always constant.
func Benchmark(b *testing.B) {
	os.Stdout = nil // Disable all program output apart from benchmark results
	b.Run("Median filter benchmark", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			filter("ship.png", "out.png")
		}
	})
}
