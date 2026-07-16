package channels

import (
	"sync"
	"time"
)

func Generate(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			ch <- i
		}
	}()
	return ch
}

func SumChannel(ch <-chan int) int {
	return 0 // TODO: range until ch closes
}

func Collect(ch <-chan int) []int {
	return nil // TODO
}

// Relay is the stepping-stone to Merge: copy one input to one output, then close.
func Relay(in <-chan int) <-chan int {
	out := make(chan int)
	close(out) // TODO: replace with a goroutine that relays and closes
	return out
}

func Merge(a, b <-chan int) <-chan int {
	out := make(chan int)
	close(out) // TODO: one copier per input; close once both copiers finish
	return out
}

// FirstResult returns whichever input produces first.
func FirstResult(a, b <-chan int) int {
	return 0 // TODO: select
}

// ReceiveWithin waits for a value or a timeout. ok is false on timeout.
func ReceiveWithin(ch <-chan int, timeout time.Duration) (value int, ok bool) {
	return 0, false // TODO: select between ch and time.After(timeout)
}

// Keep sync in the starter imports for Merge.
var _ sync.WaitGroup
