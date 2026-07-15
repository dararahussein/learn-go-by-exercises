package set04

// Test code, don't change this!

import (
	"maps"
	"sync"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	c := &SafeCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	if c.Value() != 1000 {
		t.Errorf("after 1000 concurrent Inc(), Value() == %d, want 1000", c.Value())
	}
}

func TestConcurrentWordCount(t *testing.T) {
	got := ConcurrentWordCount([]string{"go go", "gopher", "go gopher"})
	want := map[string]int{"go": 3, "gopher": 2}
	if !maps.Equal(got, want) {
		t.Errorf("ConcurrentWordCount == %v, want %v", got, want)
	}
}
