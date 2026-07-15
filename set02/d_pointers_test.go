package set02

// Test code, don't change this!

import "testing"

func TestSetZero(t *testing.T) {
	x := 42
	SetZero(&x)
	if x != 0 {
		t.Errorf("after SetZero(&x), x == %d, want 0", x)
	}
}

func TestIncrement(t *testing.T) {
	x := 5
	Increment(&x)
	if x != 6 {
		t.Errorf("after Increment(&x) on 5, x == %d, want 6", x)
	}
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Errorf("after Swap(&a, &b), a,b == %d,%d, want 2,1", a, b)
	}
}

func TestCounter(t *testing.T) {
	c := &Counter{}
	c.Inc()
	c.Inc()
	c.Inc()
	if c.Value() != 3 {
		t.Errorf("after three Inc(), Value() == %d, want 3", c.Value())
	}
}
