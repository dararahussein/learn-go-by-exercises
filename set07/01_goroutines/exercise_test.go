package goroutines

import (
	"slices"
	"sync/atomic"
	"testing"
)

func TestExercises(t *testing.T) {
	steps := []struct {
		name string
		run  func(*testing.T)
	}{
		{"01_RunAll_worked_example", testRunAll},
		{"02_Parallel", testParallel},
		{"03_SumConcurrent", testSumConcurrent},
	}
	for _, step := range steps {
		if !t.Run(step.name, step.run) {
			return
		}
	}
}

func testRunAll(t *testing.T) {
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

func testParallel(t *testing.T) {
	fns := []func() int{func() int { return 1 }, func() int { return 2 }, func() int { return 3 }}
	if got := Parallel(fns); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("Parallel = %v; want [1 2 3]", got)
	}
}

func testSumConcurrent(t *testing.T) {
	if got := SumConcurrent([][]int{{1, 2}, {3, 4}, {5}}); got != 15 {
		t.Errorf("SumConcurrent = %d; want 15", got)
	}
}
