package pointers

import "testing"

func TestExercises(t *testing.T) {
	steps := []struct {
		name string
		run  func(*testing.T)
	}{
		{"01_SetZero_worked_example", testSetZero},
		{"02_Increment", testIncrement},
		{"03_Counter", testCounter},
	}
	for _, step := range steps {
		step.run(t)
	}
}

func testSetZero(t *testing.T) {
	x := 42
	SetZero(&x)
	if x != 0 {
		t.Fatalf("after SetZero(&x), x: got %d, want 0", x)
	}
}

func testIncrement(t *testing.T) {
	x := 5
	Increment(&x)
	if x != 6 {
		t.Fatalf("after Increment(&x), x: got %d, want 6", x)
	}
}

func testCounter(t *testing.T) {
	c := &Counter{}
	c.Inc()
	c.Inc()
	if got := c.Value(); got != 2 {
		t.Fatalf("Value(): got %d, want 2", got)
	}
}
