package stats

import (
	"strconv"
	"testing"
	"time"
)

func TestStat(t *testing.T) {

	// sample data
	nums := []float64{1.0, 2.0, 3.0, 4.0, 2.0}

	list := NewList()
	counter := NewCounter(list)

	// Adds to the list aswell
	counter.AddByArray(nums)

	if !list.Equal(counter) {
		t.Errorf("List and Counter do not match.")
	}
	if list.GetSum() != counter.GetSum() {
		t.Errorf("Sum is incorrect!")
	}
	if list.GetMean() != counter.GetMean() {
		t.Errorf("Mean is incorrect!")
	}
	if counter.GetCommon() < 0 {
		t.Errorf("Mode is incorrect!")
	}

}

func TestPlot(t *testing.T) {
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}
	var m map[string]time.Duration = make(map[string]time.Duration, len(round))
	for i, v := range round {
		m[strconv.Itoa(i)] = v
	}

	if !PlotData(m) {
		t.Errorf("PlotData returned false.")
	}
}
