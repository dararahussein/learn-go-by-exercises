package set03

import (
	"errors"
	"strings"
	"testing"
)

func TestHead(t *testing.T) {
	if got, err := Head([]int{4, 5}); got != 4 || err != nil {
		t.Errorf("Head([4 5]) = (%d, %v); want (4, nil)", got, err)
	}
	if _, err := Head(nil); err == nil {
		t.Error("Head(nil) should return an error")
	}
}

func TestTail(t *testing.T) {
	got, err := Tail([]int{1, 2, 3})
	if err != nil || len(got) != 2 || got[0] != 2 {
		t.Errorf("Tail([1 2 3]) = (%v, %v); want ([2 3], nil)", got, err)
	}
	if _, err := Tail(nil); err == nil {
		t.Error("Tail(nil) should return an error")
	}
}

func TestDivFirstTwo(t *testing.T) {
	if got, err := DivFirstTwo([]int{10, 2}); got != 5 || err != nil {
		t.Errorf("DivFirstTwo([10 2]) = (%d, %v); want (5, nil)", got, err)
	}
	for _, xs := range [][]int{{10}, {10, 0}} {
		if _, err := DivFirstTwo(xs); err == nil {
			t.Errorf("DivFirstTwo(%v) should return an error", xs)
		}
	}
}

func TestLookupAgeWrapsSentinel(t *testing.T) {
	_, err := LookupAge(map[string]int{"Ada": 36}, "Grace")
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("error %v should wrap ErrNotFound", err)
	}
	if err == nil || !strings.Contains(err.Error(), "Grace") {
		t.Errorf("error %v should mention Grace", err)
	}
}

func TestTeamAgePreservesCause(t *testing.T) {
	_, err := TeamAge(map[string]int{"Ada": 36}, []string{"Ada", "Grace"})
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("TeamAge error %v should preserve ErrNotFound", err)
	}
}
