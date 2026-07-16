package channels

import (
	"slices"
	"testing"
	"time"
)

func TestExercises(t *testing.T) {
	if !t.Run("01_ChannelBasics", testChannelBasics) {
		return
	}
	if !t.Run("02_Merge", testMerge) {
		return
	}
	t.Run("03_Select", testSelect)
}

func testChannelBasics(t *testing.T) {
	if got := SumChannel(Generate(5)); got != 15 {
		t.Errorf("SumChannel\n  got:  %d\n  want: 15", got)
	}
	if got := Collect(Relay(Generate(3))); !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("Collect(Relay)\n  got:  %v\n  want: [1 2 3]", got)
	}
}

func testMerge(t *testing.T) {
	total := 0
	for value := range Merge(Generate(3), Generate(3)) {
		total += value
	}
	if total != 12 {
		t.Errorf("merged total\n  got:  %d\n  want: 12", total)
	}
}

func testSelect(t *testing.T) {
	slow := make(chan int, 1)
	fast := make(chan int, 1)
	fast <- 7
	if got := FirstResult(slow, fast); got != 7 {
		t.Errorf("FirstResult\n  got:  %d\n  want: 7", got)
	}
	if _, ok := ReceiveWithin(make(chan int), 5*time.Millisecond); ok {
		t.Error("ReceiveWithin timeout flag\n  got:  true\n  want: false")
	}
}

/*
Q: Would SumConcurrent beat a plain loop for three slices containing five
   integers each? Account for goroutine/channel overhead in your answer.
A:
*/
