package errorx

import (
	"errors"
	"strings"
	"testing"
)

func TestExercises(t *testing.T) {
	steps := []struct {
		name string
		run  func(*testing.T)
	}{
		{"01_Head_worked_example", testHead},
		{"02_Tail", testTail},
		{"03_DivFirstTwo", testDivFirstTwo},
		{"04_LookupAge", testLookupAgeWrapsSentinel},
		{"05_TeamAge", testTeamAgePreservesCause},
	}
	for _, step := range steps {
		step.run(t)
	}
}

func testHead(t *testing.T) {
	if got, err := Head([]int{4, 5}); got != 4 || err != nil {
		t.Fatalf("Head([4 5]): got (%d, %v), want (4, nil)", got, err)
	}
	if _, err := Head(nil); err == nil {
		t.Fatal("Head(nil) error: got nil, want non-nil")
	}
}

func testTail(t *testing.T) {
	got, err := Tail([]int{1, 2, 3})
	if err != nil || len(got) != 2 || got[0] != 2 {
		t.Fatalf("Tail([1 2 3]): got (%v, %v), want ([2 3], nil)", got, err)
	}
	if _, err := Tail(nil); err == nil {
		t.Fatal("Tail(nil) error: got nil, want non-nil")
	}
}

func testDivFirstTwo(t *testing.T) {
	if got, err := DivFirstTwo([]int{10, 2}); got != 5 || err != nil {
		t.Fatalf("DivFirstTwo([10 2]): got (%d, %v), want (5, nil)", got, err)
	}
	for _, xs := range [][]int{{10}, {10, 0}} {
		if _, err := DivFirstTwo(xs); err == nil {
			t.Fatalf("DivFirstTwo(%v) error: got nil, want non-nil", xs)
		}
	}
}

func testLookupAgeWrapsSentinel(t *testing.T) {
	_, err := LookupAge(map[string]int{"Ada": 36}, "Grace")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("LookupAge error: got %v, want error wrapping ErrNotFound", err)
	}
	if err == nil || !strings.Contains(err.Error(), "Grace") {
		t.Fatalf("LookupAge error text: got %v, want text mentioning Grace", err)
	}
}

func testTeamAgePreservesCause(t *testing.T) {
	_, err := TeamAge(map[string]int{"Ada": 36}, []string{"Ada", "Grace"})
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("TeamAge error: got %v, want error wrapping ErrNotFound", err)
	}
}
