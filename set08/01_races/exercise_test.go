package races

import (
	"flag"
	"maps"
	"sync"
	"testing"
)

var runRaceDemo = flag.Bool("race-demo", false, "run the deliberately racy counter test")

func TestExercises(t *testing.T) {
	if !t.Run("01_ObserveUnsafeCounter", testUnsafeCounterRace) {
		return
	}
	if !t.Run("02_SafeCounter", testSafeCounter) {
		return
	}
	t.Run("03_ConcurrentWordCount", testConcurrentWordCount)
}

func testUnsafeCounterRace(t *testing.T) {
	if !*runRaceDemo {
		t.Skip("deliberate race demo; run with -args -race-demo")
	}
	var c UnsafeCounter
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); c.Inc() }()
	}
	wg.Wait()
	_ = c.Value() // the race detector, not this assertion, is the teacher
}

func testSafeCounter(t *testing.T) {
	var c SafeCounter
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); c.Inc() }()
	}
	wg.Wait()
	if got := c.Value(); got != 1000 {
		t.Errorf("Value\n  got:  %d\n  want: 1000", got)
	}
}

func testConcurrentWordCount(t *testing.T) {
	got := ConcurrentWordCount([]string{"go go", "gopher", "go gopher"})
	want := map[string]int{"go": 3, "gopher": 2}
	if !maps.Equal(got, want) {
		t.Errorf("ConcurrentWordCount\n  got:  %v\n  want: %v", got, want)
	}
}
