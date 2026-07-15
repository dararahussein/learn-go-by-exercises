package set04

import "testing"

func TestNormalizeSpaces(t *testing.T) {
	// This passes even while NormalizeSpaces is incorrect. Add at least four
	// cases that cover trimming, three spaces, tabs, newlines, and empty input.
	if got := NormalizeSpaces("hello  go"); got != "hello go" {
		t.Errorf("NormalizeSpaces = %q; want %q", got, "hello go")
	}
}

/*
REFLECTION (answer in this file):

Q1: Why did the original test pass for an incorrect implementation?
A1:
Q2: What does `go test -run TestNormalizeSpaces ./set04/` target?
A2:
Q3: What did `go test -cover` reveal, and what does coverage not prove?
A3:
*/
