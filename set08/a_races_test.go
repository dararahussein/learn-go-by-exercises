package set08

import (
	"maps"
	"sync"
	"testing"
)

func TestUnsafeCounterRace(t *testing.T) {
	var c UnsafeCounter
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); c.Inc() }()
	}
	wg.Wait()
	_ = c.Value() // the race detector, not this assertion, is the teacher
}

func TestSafeCounter(t *testing.T) {
	var c SafeCounter
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); c.Inc() }()
	}
	wg.Wait()
	if got := c.Value(); got != 1000 {
		t.Errorf("Value = %d; want 1000", got)
	}
}

func TestConcurrentWordCount(t *testing.T) {
	got := ConcurrentWordCount([]string{"go go", "gopher", "go gopher"})
	want := map[string]int{"go": 3, "gopher": 2}
	if !maps.Equal(got, want) {
		t.Errorf("ConcurrentWordCount = %v; want %v", got, want)
	}
}
