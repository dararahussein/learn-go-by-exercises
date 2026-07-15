package set07

import (
	"slices"
	"testing"
	"time"
)

func TestChannelBasics(t *testing.T) {
	if got := SumChannel(Generate(5)); got != 15 {
		t.Errorf("SumChannel = %d; want 15", got)
	}
	if got := Collect(Relay(Generate(3))); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("Collect(Relay) = %v; want [1 2 3]", got)
	}
}

func TestMerge(t *testing.T) {
	total := 0
	for value := range Merge(Generate(3), Generate(3)) {
		total += value
	}
	if total != 12 {
		t.Errorf("merged total = %d; want 12", total)
	}
}

func TestSelect(t *testing.T) {
	slow := make(chan int, 1)
	fast := make(chan int, 1)
	fast <- 7
	if got := FirstResult(slow, fast); got != 7 {
		t.Errorf("FirstResult = %d; want 7", got)
	}
	if _, ok := ReceiveWithin(make(chan int), 5*time.Millisecond); ok {
		t.Error("ReceiveWithin should time out")
	}
}

/*
Q: Would SumConcurrent beat a plain loop for three slices containing five
   integers each? Account for goroutine/channel overhead in your answer.
A:
*/
