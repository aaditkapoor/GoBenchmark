// Package benchmark provides the representation of a single benchmark consisting
// of a function and its converted types.
package benchmark // import aaditkapoor/GoBenchmark/models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
)

// ConversionType represents a unit that can be converted.
type ConversionType string

const (
	// Milli Milliseconds
	Milli ConversionType = "ms"
	// Micro Microseconds
	Micro = "us"
	// Nano Nanoseconds
	Nano = "ns"
	// Sec Seconds
	Sec = "s"
	// Min Minutes
	Min = "m"
	// Hr Hours
	Hr = "hr"
)

// Benchmark Represents a single instance of a benchmark.
type Benchmark struct {
	description    string
	function       func()
	currentTime    time.Time
	elapsed        time.Duration
	convertedTypes map[ConversionType]string
	iterations     int
	units          []ConversionType
	states         map[string]time.Duration
}

// Start starts the timer for measuring the function.
func (benchmark *Benchmark) Start() {
	benchmark.currentTime = time.Now()
}

// GetStates get time.Duration at each time in a key value pair of iterations and durations.
func (benchmark *Benchmark) GetStates() map[string]time.Duration {
	return benchmark.states
}

// GetElapsedTime get the the elapsed time.
func (benchmark Benchmark) GetElapsedTime() time.Duration {
	return benchmark.elapsed
}

// GetUnits get the units that the user wants during execution.
func (benchmark Benchmark) GetUnits() []ConversionType {
	return benchmark.units
}

// GetFunction get the measuring function.
func (benchmark Benchmark) GetFunction() func() {
	return benchmark.function
}

// GetIterations get the number of iterations.
func (benchmark Benchmark) GetIterations() int {
	return benchmark.iterations
}

// GetConvertedType get ConvertedType according to t (ConversionType).
func (benchmark *Benchmark) GetConvertedType(t ConversionType) time.Duration {
	duration, _ := time.ParseDuration(benchmark.elapsed.String())
	var converted time.Duration
	switch t {
	case Milli:
		converted = duration.Round(time.Millisecond)
	case Micro:
		converted = duration.Round(time.Microsecond)
	case Nano:
		converted = duration.Round(time.Nanosecond)
	case Sec:
		converted = duration.Round(time.Second)
	case Min:
		converted = duration.Round(time.Minute)
	case Hr:
		converted = duration.Round(time.Hour)
	}

	return converted
}

// TimeElapsed calculates time elapsed and performs conversion of units.
func (benchmark *Benchmark) TimeElapsed(units []ConversionType) {

	benchmark.elapsed = time.Since(benchmark.currentTime)
	benchmark.units = units

	// Parse the duration (time.Duration)
	duration, _ := time.ParseDuration(benchmark.elapsed.String())

	if len(units) == 0 {
		fmt.Println("Using default conversion type: Microseconds.")
		benchmark.convertedTypes[Micro] = fmt.Sprintf("%d microseconds", duration.Microseconds())
	}

	for _, ctype := range units {

		switch ctype {

		case Milli:
			benchmark.convertedTypes[Milli] = fmt.Sprintf("%d milliseconds", duration.Milliseconds())
		case Micro:
			benchmark.convertedTypes[Micro] = fmt.Sprintf("%d microseconds", duration.Microseconds())
		case Nano:
			benchmark.convertedTypes[Nano] = fmt.Sprintf("%d nanoseconds", duration.Nanoseconds())
		case Sec:
			benchmark.convertedTypes[Sec] = fmt.Sprintf("%f seconds", duration.Seconds())
		case Min:
			benchmark.convertedTypes[Min] = fmt.Sprintf("%f minutes", duration.Minutes())
		case Hr:
			benchmark.convertedTypes[Hr] = fmt.Sprintf("%f hours", duration.Hours())
		default:
			benchmark.convertedTypes[Micro] = fmt.Sprintf("%d microseconds", duration.Microseconds())
		}
	}
}

// Main start running the benchmark.
func (benchmark Benchmark) Main() {
	timeTaken := ""
	for _, val := range benchmark.convertedTypes {
		spf := fmt.Sprintf("(%s) ", val)
		timeTaken += spf
	}

	s := fmt.Sprintf("Running %s ...[%s] %s with [%s] iterations.", color.GreenString(benchmark.description), color.GreenString("done"), color.RedString(timeTaken), color.RedString(strconv.Itoa(benchmark.iterations)))
	fmt.Println(s)
}

// NewBenchmark create a new Benchmark Object.
func NewBenchmark(description string, function func(), iterations int, units ...ConversionType) Benchmark {
	benchmark := Benchmark{description: description, function: function, iterations: iterations, convertedTypes: make(map[ConversionType]string, len(units))}

	var m map[string]time.Duration = make(map[string]time.Duration)
	fmt.Printf("======%s======\n", GetFunctionName(function))
	for i := 0; i < iterations+1; i++ {
		benchmark.Start()
		function()
		benchmark.TimeElapsed(units)
		m[strconv.Itoa(i)] = benchmark.elapsed
	}
	benchmark.states = m
	fmt.Printf("======%s======\n", GetFunctionName(function))
	return benchmark
}
