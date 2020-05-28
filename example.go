// A glimpse of using the package.
package main

import (
	"fmt"
	"math/rand"

	"github.com/aaditkapoor/GoBenchmark/benchmark"
	"github.com/aaditkapoor/GoBenchmark/stats"
)

// Demo Function
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func main() {

	// --------- STATISTICS ----------

	// Creating a benchmark object with 5 iterations
	b := benchmark.NewBenchmark("measuring quicksort algorithm", func() {
		n := quicksort([]int{2312, 212, 2, 31, 33, 0, 23})
		fmt.Println(n)

	}, 5, benchmark.Micro, benchmark.Nano) // more options: see Benchmark.go

	// Creating a benchmark stat object
	bs := stats.NewBenchmarkStat(b, stats.All) // more options: see stats.StatType

	// Printing stat
	bs.PrintStats()

	// ----- END -------------

	// We can also run a plain benchmark by calling .Main()

	//b.Main()
}
