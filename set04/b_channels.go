/*

Channels — how goroutines talk to each other.

A channel is a typed pipe. One goroutine sends, another receives:

	ch := make(chan int)
	ch <- 7        // send (blocks until someone receives)
	v := <-ch      // receive
	close(ch)      // no more values will be sent
	for v := range ch { ... }   // receive until the channel is closed

`<-chan int` in a signature means "a receive-only channel" — the caller can read
but not send. It documents direction and the compiler enforces it.

Tasks:
- [ ] Implement the stubs (marked TODO).
- [ ] Run `go test ./set04/` after every change.

Expected time to complete: 45 minutes.

*/

package set04

// DONE FOR YOU: send 1..n on a channel from a background goroutine, then close.
// Returning the channel immediately lets the caller start receiving while this
// goroutine is still producing.
func Generate(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch) // tells `range` on the other end to stop
	}()
	return ch
}

// SumChannel receives every value from ch (until it's closed) and returns the
// total. `for v := range ch` is the idiom.
func SumChannel(ch <-chan int) int {
	return 0 // TODO: implement SumChannel
}

// Collect drains ch into a slice, in the order values arrive.
func Collect(ch <-chan int) []int {
	return nil // TODO: implement Collect
}

// Merge fans two channels into one: every value from a AND b comes out the
// returned channel, which is closed once both inputs are drained. Output order
// doesn't matter. A clean way: launch a goroutine per input that copies values
// into out (coordinate with a WaitGroup), plus one goroutine that Wait()s and
// then close(out)s.
func Merge(a, b <-chan int) <-chan int {
	out := make(chan int)
	close(out) // TODO: replace this — right now it returns an empty channel
	return out
}
