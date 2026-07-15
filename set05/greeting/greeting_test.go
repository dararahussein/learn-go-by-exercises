package greeting

// Test code, don't change this!

import "testing"

func TestGreet(t *testing.T) {
	cases := []struct{ name, want string }{
		{"Ada", "Hello, Ada! Welcome to Go."},
		{"", "Hello, friend! Welcome to Go."},
	}
	for _, c := range cases {
		if got := Greet(c.name); got != c.want {
			t.Errorf("Greet(%q) == %q, want %q", c.name, got, c.want)
		}
	}
}
