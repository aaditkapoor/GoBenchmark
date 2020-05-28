// Package stats Utilities that help calculating statistical measures.
package stats

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/wcharczuk/go-chart"
)

// List base list type
type List struct {
	capacity int
	items    []float64
}

// NewList create a new list.
func NewList() *List {
	l := List{}
	return &l
}

// Add add item to the list.
func (list *List) Add(e float64) {
	list.items = append(list.items, e)
}

// String string representation of a list.
func (list List) String() string {
	return fmt.Sprintf("%v", list.items)
}

// GetSum return the sum of items.
func (list *List) GetSum() float64 {
	sum := 0.0
	for _, val := range list.items[1:] {
		sum += val
	}

	return sum
}

// Equal Check whether two interfaces (Only List and Counter)
// are equal.
func (list *List) Equal(other interface{}) bool {
	// Type asserting to only allow Counter types
	if other, ok := other.(Counter); ok {
		return cmp.Equal(list.items, other.items)
	}
	return false

}

// return the min and max of an array.
func minMax(array []float64) (float64, float64) {
	var max float64 = array[0]
	var min float64 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

// GetMean return the mean of items.
func (list *List) GetMean() float64 {
	return list.GetSum() / float64(len(list.items))
}

// GetLowest return the lowest value.
func (list *List) GetLowest() float64 {
	min, _ := minMax(list.items)
	return min
}

// GetHighest return the higest value.
func (list *List) GetHighest() float64 {
	_, max := minMax(list.items)
	return max
}

// Counter A type composed from List capable of keeping track of each element in a key value pair.
type Counter struct {
	*List
	mapped map[float64]int
}

// NewCounter creates a new Counter.
func NewCounter(list *List) Counter {
	c := Counter{List: list, mapped: make(map[float64]int, len(list.items))}
	return c
}

// Add add to counter, if element exists then increment count.
// Does this in a separate map and also adds to the main item array.
func (counter *Counter) Add(e float64) {
	// e is already there
	if _, ok := counter.mapped[e]; ok {
		counter.mapped[e]++
	} else {
		counter.mapped[e] = 1
	}
}

// AddByArray add to the Counter and List by an array.
// OPTION() Can remove the abstraction.
func (counter *Counter) AddByArray(arr []float64) {
	for _, val := range arr {
		counter.Add(val)
		counter.List.Add(val) // Add to the list aswell
	}
}

// GetCommon returns the mode of the data.
func (counter *Counter) GetCommon() float64 {
	var keys []float64
	var vals []int
	for k, v := range counter.mapped {
		keys = append(keys, k)
		vals = append(vals, v)
	}

	// get max value
	max := vals[0]
	for _, v := range vals {
		if max < v {
			max = v
		}
	}

	// get key
	key := 0.0
	for k, v := range counter.mapped {
		if v == max {
			key = k
		}
	}

	// returns key of the highest mode
	return key

}

// PlotData plots the data given a map of iterations and durations.
// BUG(r) Function does not work as expected.
func PlotData(data map[string]time.Duration) bool {

	var keys []float64
	var values []float64

	for key, val := range data {

		// Add float values
		if s, err := strconv.ParseFloat(key, 32); err == nil {
			keys = append(keys, s)
		}

		// Add time values
		values = append(values, float64(val.Microseconds()))
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "Iterations",
		},
		YAxis: chart.YAxis{
			Name: "Time (μs)",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: keys,
				YValues: values,
			},
		},
	}

	f, _ := os.Create("graph.png")
	defer f.Close()
	err := graph.Render(chart.PNG, f)
	if err == nil {
		return true
	}
	fmt.Println(err)
	return false

}
