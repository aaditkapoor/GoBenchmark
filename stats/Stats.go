// Package stats Provides the StatBenchmark type that is capable of getting statistics
// for a benchmark given number of iterations.
package stats

import (
	"fmt"
	"time"

	bmark "github.com/aaditkapoor/GoBenchmark"
	"github.com/fatih/color"
)

// StatType represents a statistical measure.
type StatType string

const (

	// Mean opt
	Mean StatType = "mean"
	// Sum opt
	Sum = "sum"
	// Mode opt
	Mode = "mode"
	// Highest opt
	Highest = "highest"
	// Lowest opt
	Lowest = "lowest"

	// Range opt
	Range = "range"

	// All opt
	All = "all"

	// Graph opt
	Graph = "graph"
)

// BenchmarkStat represents a statistical Benchmark
type BenchmarkStat struct {
	benchmark bmark.Benchmark
	mean      float64
	sum       float64
	mode      float64
	frange    float64
	highest   float64
	lowest    float64
	options   []StatType
}

// Set the mean.
func (benchmark *BenchmarkStat) setMean(mean float64) {
	benchmark.mean = mean
}

// Set the range.
func (benchmark *BenchmarkStat) setRange(float64) {
	benchmark.frange = benchmark.highest - benchmark.lowest
}

// Set the sum.
func (benchmark *BenchmarkStat) setSum(sum float64) {
	benchmark.sum = sum
}

// Set the mode.
func (benchmark *BenchmarkStat) setMode(mode float64) {
	benchmark.mode = mode
}

// Set the highest time.
func (benchmark *BenchmarkStat) setHighest(highest float64) {
	benchmark.highest = highest
}

// Set the lowest time.
func (benchmark *BenchmarkStat) setLowest(lowest float64) {
	benchmark.lowest = lowest
}

// GetStat return a computed StatType
func (benchmark BenchmarkStat) GetStat(what StatType) (string, error) {

	switch what {
	case Mean:
		return fmt.Sprintf("%s = %fμs\n", color.GreenString(string(what)), benchmark.mean), nil
	case Mode:
		return fmt.Sprintf("%s = %fμs\n", color.GreenString(string(what)), benchmark.mode), nil
	case Highest:
		return fmt.Sprintf("%s = %fμs\n", color.GreenString(string(what)), benchmark.highest), nil
	case Lowest:
		return fmt.Sprintf("%s = %fμs\n", color.GreenString(string(what)), benchmark.lowest), nil
	case Sum:
		return fmt.Sprintf("%s = %fμs\n", color.GreenString(string(what)), benchmark.sum), nil
	case Range:
		return fmt.Sprintf("%s = %fμs\n", color.GreenString(string(what)), benchmark.frange), nil
	case All:
		return fmt.Sprintf("%s=%fμs\n%s=%fμs\n%s=%fμs\n%s=%fμs\n%s=%fμs\n", color.GreenString("mean"), benchmark.mean, color.GreenString("mode"), benchmark.mode, color.GreenString("highest"),
			benchmark.highest, color.GreenString("lowest"), benchmark.lowest, color.GreenString("sum"), benchmark.sum), nil
	default:
		return fmt.Sprintf("Invalid type: %s", what), fmt.Errorf(color.RedString("Invalid type: %s", what))
	}
}

// Parse options from NewBenchmark.
func parseOptions(benchmark *BenchmarkStat, options []StatType, m map[string]time.Duration) {

	var iterations []string
	var times []float64
	list := NewList()
	counter := NewCounter(list)

	for k, v := range m {
		iterations = append(iterations, k)
		if d, err := time.ParseDuration(v.String()); err == nil {
			times = append(times, float64(d.Microseconds()))
		}
	}

	counter.AddByArray(times)

	for _, val := range options {
		switch val {
		case Mean:
			benchmark.setMean(counter.GetMean())
		case Sum:
			benchmark.setSum(counter.GetSum())
		case Mode:
			benchmark.setMode(counter.GetCommon())
		case Lowest:
			benchmark.setLowest(counter.GetLowest())
		case Highest:
			benchmark.setHighest(counter.GetHighest())
		case Range:
			benchmark.setRange(counter.GetHighest() - counter.GetLowest())
		case Graph:
			if !PlotData(m) {
				fmt.Println(color.RedString("Error in plotting the graph."))
			} else {
				fmt.Println(color.GreenString("Graph saved as [graph.png]"))
			}
		case All:
			benchmark.setMean(counter.GetMean())
			benchmark.setSum(counter.GetSum())
			benchmark.setMode(counter.GetCommon())
			benchmark.setLowest(counter.GetLowest())
			benchmark.setHighest(counter.GetHighest())
			benchmark.setRange(counter.GetHighest() - counter.GetLowest())
		}
	}
}

// PrintStats prints the selected stats options.
func (benchmark *BenchmarkStat) PrintStats() {
	printable := ""
	printable += "\nSTATISTICS\n"
	printable += "===============\n"
	for _, val := range benchmark.options {
		switch val {
		case Mean:
			if st, err := benchmark.GetStat(val); err == nil {
				printable += fmt.Sprintf("%s", st)
			}
		case Mode:
			if st, err := benchmark.GetStat(val); err == nil {
				printable += fmt.Sprintf("%s", st)
			}
		case Highest:
			if st, err := benchmark.GetStat(val); err == nil {
				printable += fmt.Sprintf("%s", st)
			}
		case Lowest:
			if st, err := benchmark.GetStat(val); err == nil {
				printable += fmt.Sprintf("%s", st)
			}
		case Sum:
			if st, err := benchmark.GetStat(val); err == nil {
				printable += fmt.Sprintf("%s", st)
			}
		case Range:
			if st, err := benchmark.GetStat(val); err == nil {
				printable += fmt.Sprintf("%s", st)
			}
		case All:
			if st, err := benchmark.GetStat(val); err == nil {
				printable += fmt.Sprintf("%s", st)
			}
		default:
			if _, err := benchmark.GetStat(val); err != nil {
				printable += err.Error()
			}

		}
	}

	s := fmt.Sprintf("%s", printable)
	fmt.Println(s)
}

// NewBenchmarkStat create a new Benchmark with statistics qualities
// NewBenchmarkStat(benchmark, "mean", "mode") || NewBenchmarkStat(benchmark, stats.Mean, stats.Mode)
func NewBenchmarkStat(benchmark bmark.Benchmark, options ...StatType) BenchmarkStat {
	bs := BenchmarkStat{benchmark: benchmark}
	iterations := benchmark.GetIterations()
	units := benchmark.GetUnits()
	function := benchmark.GetFunction()
	fmt.Printf("\nDescription: %s\n", color.GreenString(benchmark.GetDesc()))
	fmt.Printf("Statistics for %s with %d iterations and %v units.\n", bmark.GetFunctionName(function), iterations, units)

	var m map[string]time.Duration = benchmark.GetStates()
	parseOptions(&bs, options, m)

	// Start the main benchmark
	//benchmark.Main()
	bs.options = options
	return bs
}
