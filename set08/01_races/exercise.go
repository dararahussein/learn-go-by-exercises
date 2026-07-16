package races

import "sync"

// UnsafeCounter is deliberately unsafe for the race-detector lab.
type UnsafeCounter struct{ n int }

func (c *UnsafeCounter) Inc()       { c.n++ } // BUG: read-modify-write is not atomic
func (c *UnsafeCounter) Value() int { return c.n }

type SafeCounter struct {
	mu sync.Mutex
	n  int
}

func (c *SafeCounter) Inc() {
	// TODO: lock, update, unlock
}

func (c *SafeCounter) Value() int {
	return 0 // TODO: protect reads with the same mutex
}

func ConcurrentWordCount(texts []string) map[string]int {
	return nil // TODO: one goroutine per text; protect the shared result map
}
