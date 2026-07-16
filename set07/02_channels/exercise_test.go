package channels

import (
	"slices"
	"testing"
	"time"
)

func TestExercises(t *testing.T) {
	testChannelBasics(t)
	testMerge(t)
	testSelect(t)
}

func testChannelBasics(t *testing.T) {
	if got := SumChannel(Generate(5)); got != 15 {
		t.Fatalf("SumChannel: got %d, want 15", got)
	}
	if got := Collect(Relay(Generate(3))); !slices.Equal(got, []int{1, 2, 3}) {
		t.Fatalf("Collect(Relay): got %v, want [1 2 3]", got)
	}
}

func testMerge(t *testing.T) {
	total := 0
	for value := range Merge(Generate(3), Generate(3)) {
		total += value
	}
	if total != 12 {
		t.Fatalf("merged total: got %d, want 12", total)
	}
}

func testSelect(t *testing.T) {
	slow := make(chan int, 1)
	fast := make(chan int, 1)
	fast <- 7
	if got := FirstResult(slow, fast); got != 7 {
		t.Fatalf("FirstResult: got %d, want 7", got)
	}
	if _, ok := ReceiveWithin(make(chan int), 5*time.Millisecond); ok {
		t.Fatal("ReceiveWithin timeout flag: got true, want false")
	}
}

/*
Q: Would SumConcurrent beat a plain loop for three slices containing five
   integers each? Account for goroutine/channel overhead in your answer.
A:
*/
