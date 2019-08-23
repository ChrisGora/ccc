package main

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func random(n int) []int32 {
	s := make([]int32, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Int31()
	}
	return s
}

const (
	start = 512
	end   = 16384
)

func BenchmarkSequential(b *testing.B) {
	for size := start; size <= end; size *= 2 {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			os.Stdout = nil // Disable all program output apart from benchmark results
			for i := 0; i < b.N; i++ {
				unsorted := random(size)
				b.StartTimer()
				mergeSort(unsorted)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkParallel(b *testing.B) {
	for size := start; size <= end; size *= 2 {
		b.Run(fmt.Sprint(size), func(b *testing.B) {
			os.Stdout = nil // Disable all program output apart from benchmark results
			for i := 0; i < b.N; i++ {
				unsorted := random(size)
				b.StartTimer()
				parallelMergeSort(unsorted)
				b.StopTimer()
			}
		})
	}
}
