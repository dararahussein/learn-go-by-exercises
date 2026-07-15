package set04

// Test code, don't change this!

import (
	"slices"
	"sync"
	"testing"
)

func TestRunAll(t *testing.T) {
	var mu sync.Mutex
	count := 0
	fns := make([]func(), 100)
	for i := range fns {
		fns[i] = func() {
			mu.Lock()
			count++
			mu.Unlock()
		}
	}
	RunAll(fns)
	if count != 100 {
		t.Errorf("RunAll ran %d of 100 functions", count)
	}
}

func TestParallel(t *testing.T) {
	fns := []func() int{
		func() int { return 1 },
		func() int { return 2 },
		func() int { return 3 },
	}
	got := Parallel(fns)
	if !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("Parallel == %v, want [1 2 3] (order must match input)", got)
	}
}

func TestSumConcurrent(t *testing.T) {
	got := SumConcurrent([][]int{{1, 2}, {3, 4}, {5}})
	if got != 15 {
		t.Errorf("SumConcurrent == %d, want 15", got)
	}
	if got := SumConcurrent(nil); got != 0 {
		t.Errorf("SumConcurrent(nil) == %d, want 0", got)
	}
}
