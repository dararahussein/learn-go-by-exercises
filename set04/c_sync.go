/*

Mutexes — protecting shared data from data races.

When several goroutines touch the SAME variable and at least one writes, you
have a data race: the result is undefined and the program is buggy. A
sync.Mutex fixes it by letting only one goroutine into the critical section:

	mu.Lock()
	// ... touch the shared data ...
	mu.Unlock()

Prove races to yourself with the race detector:  go test -race ./set04/
Implement Inc WITHOUT the lock first, run with -race, and watch it complain.

Tasks:
- [ ] Implement the stubs (marked TODO).
- [ ] Run `go test ./set04/` (and try `go test -race ./set04/`).

Expected time to complete: 45 minutes.

*/

package set04

import "sync"

// A counter that's safe to hammer from many goroutines at once.
type SafeCounter struct {
	mu sync.Mutex
	n  int
}

// DONE FOR YOU: read the value under the lock.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

// Inc adds 1 to the counter. Take the lock, change c.n, release the lock.
// (defer c.mu.Unlock() right after Lock() is the usual, leak-proof style.)
func (c *SafeCounter) Inc() {
	// TODO: implement Inc (safely!)
}

// ConcurrentWordCount counts words across many texts, processing each text in
// its own goroutine and merging the counts. The merged map is SHARED, so every
// update to it must be under a mutex (or you'll get a race and wrong totals).
// strings.Fields(text) splits a string into words — add the "strings" import.
// Result is the same regardless of goroutine order.
func ConcurrentWordCount(texts []string) map[string]int {
	return nil // TODO: implement ConcurrentWordCount
}
