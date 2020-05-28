package main

import (
	"aaditkapoor/GoBenchmark/benchmark"
	"aaditkapoor/GoBenchmark/stats"
	"fmt"
)

// demo function
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n + factorial(n-1)
}

func main() {

	// Creating a benchmark object with 5 iterations
	b := benchmark.NewBenchmark("factorial program", func() {
		n := factorial(13)
		fmt.Println(n)

	}, 5, benchmark.Micro, benchmark.Nano)

	// Creating a benchmark stat object
	bs := stats.NewBenchmarkStat(b, stats.Mean)

	// Printing stat
	bs.PrintStats()
}
