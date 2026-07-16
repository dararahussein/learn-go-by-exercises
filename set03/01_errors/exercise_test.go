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
		if !t.Run(step.name, step.run) {
			return
		}
	}
}

func testHead(t *testing.T) {
	if got, err := Head([]int{4, 5}); got != 4 || err != nil {
		t.Errorf("Head([4 5])\n  got:  (%d, %v)\n  want: (4, nil)", got, err)
	}
	if _, err := Head(nil); err == nil {
		t.Error("Head(nil) error\n  got:  nil\n  want: non-nil")
	}
}

func testTail(t *testing.T) {
	got, err := Tail([]int{1, 2, 3})
	if err != nil || len(got) != 2 || got[0] != 2 {
		t.Errorf("Tail([1 2 3])\n  got:  (%v, %v)\n  want: ([2 3], nil)", got, err)
	}
	if _, err := Tail(nil); err == nil {
		t.Error("Tail(nil) error\n  got:  nil\n  want: non-nil")
	}
}

func testDivFirstTwo(t *testing.T) {
	if got, err := DivFirstTwo([]int{10, 2}); got != 5 || err != nil {
		t.Errorf("DivFirstTwo([10 2])\n  got:  (%d, %v)\n  want: (5, nil)", got, err)
	}
	for _, xs := range [][]int{{10}, {10, 0}} {
		if _, err := DivFirstTwo(xs); err == nil {
			t.Errorf("DivFirstTwo(%v) error\n  got:  nil\n  want: non-nil", xs)
		}
	}
}

func testLookupAgeWrapsSentinel(t *testing.T) {
	_, err := LookupAge(map[string]int{"Ada": 36}, "Grace")
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("LookupAge error\n  got:  %v\n  want: error wrapping ErrNotFound", err)
	}
	if err == nil || !strings.Contains(err.Error(), "Grace") {
		t.Errorf("LookupAge error text\n  got:  %v\n  want: text mentioning Grace", err)
	}
}

func testTeamAgePreservesCause(t *testing.T) {
	_, err := TeamAge(map[string]int{"Ada": 36}, []string{"Ada", "Grace"})
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("TeamAge error\n  got:  %v\n  want: error wrapping ErrNotFound", err)
	}
}
