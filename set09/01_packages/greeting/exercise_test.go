package greeting

import "testing"

func TestGreet(t *testing.T) {
	for name, want := range map[string]string{"Ada": "Hello, Ada!", "": "Hello, friend!"} {
		if got := Greet(name); got != want {
			t.Fatalf("Greet(%q): got %q, want %q", name, got, want)
		}
	}
}
