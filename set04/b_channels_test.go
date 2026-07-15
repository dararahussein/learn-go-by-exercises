package set04

// Test code, don't change this!

import (
	"slices"
	"testing"
)

func TestSumChannel(t *testing.T) {
	if got := SumChannel(Generate(5)); got != 15 {
		t.Errorf("SumChannel(Generate(5)) == %d, want 15", got)
	}
}

func TestCollect(t *testing.T) {
	got := Collect(Generate(3))
	if !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("Collect(Generate(3)) == %v, want [1 2 3]", got)
	}
}

func TestMerge(t *testing.T) {
	// Drain the merged channel directly and sum — order-independent, so this
	// doesn't depend on how Merge interleaves.
	sum := 0
	for v := range Merge(Generate(3), Generate(3)) {
		sum += v
	}
	if sum != 12 { // (1+2+3) + (1+2+3)
		t.Errorf("sum of Merge(Generate(3), Generate(3)) == %d, want 12", sum)
	}
}
