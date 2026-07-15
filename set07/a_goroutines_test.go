package set07

import (
	"slices"
	"sync/atomic"
	"testing"
)

func TestRunAll(t *testing.T) {
	var count atomic.Int64
	fns := make([]func(), 100)
	for i := range fns {
		fns[i] = func() { count.Add(1) }
	}
	RunAll(fns)
	if got := count.Load(); got != 100 {
		t.Errorf("RunAll count = %d; want 100", got)
	}
}

func TestParallel(t *testing.T) {
	fns := []func() int{func() int { return 1 }, func() int { return 2 }, func() int { return 3 }}
	if got := Parallel(fns); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("Parallel = %v; want [1 2 3]", got)
	}
}

func TestSumConcurrent(t *testing.T) {
	if got := SumConcurrent([][]int{{1, 2}, {3, 4}, {5}}); got != 15 {
		t.Errorf("SumConcurrent = %d; want 15", got)
	}
}
