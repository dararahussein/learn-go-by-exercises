package set02

// Test code, don't change this!

import (
	"slices"
	"strings"
	"testing"
)

func TestHead(t *testing.T) {
	if v, err := Head([]int{1, 2, 3}); v != 1 || err != nil {
		t.Errorf("Head([1 2 3]) == (%d, %v), want (1, nil)", v, err)
	}
	if v, err := Head(nil); err == nil {
		t.Errorf("Head(nil) == (%d, nil), want an error", v)
	}
}

func TestTail(t *testing.T) {
	got, err := Tail([]int{1, 2, 3})
	if err != nil || !slices.Equal(got, []int{2, 3}) {
		t.Errorf("Tail([1 2 3]) == (%v, %v), want ([2 3], nil)", got, err)
	}
	if got, err := Tail(nil); err == nil {
		t.Errorf("Tail(nil) == (%v, nil), want an error", got)
	}
}

func TestDiv(t *testing.T) {
	if v, err := Div(10, 2); v != 5 || err != nil {
		t.Errorf("Div(10, 2) == (%d, %v), want (5, nil)", v, err)
	}
	if v, err := Div(10, 0); err == nil {
		t.Errorf("Div(10, 0) == (%d, nil), want an error", v)
	}
	if v, _ := Div(10, 0); v != 0 {
		t.Errorf("Div(10, 0) returned value %d with its error; return the zero value 0", v)
	}
}

func TestDivFirstTwo(t *testing.T) {
	if v, err := DivFirstTwo([]int{10, 2, 3}); v != 5 || err != nil {
		t.Errorf("DivFirstTwo([10 2 3]) == (%d, %v), want (5, nil)", v, err)
	}
	if _, err := DivFirstTwo([]int{10, 0, 3}); err == nil {
		t.Error("DivFirstTwo([10 0 3]) should propagate the divide-by-zero error")
	}
	if _, err := DivFirstTwo([]int{10}); err == nil {
		t.Error("DivFirstTwo([10]) should error: not enough elements")
	}
	if _, err := DivFirstTwo(nil); err == nil {
		t.Error("DivFirstTwo(nil) should error: not enough elements")
	}
}

func TestFindEven(t *testing.T) {
	if v, err := FindEven([]int{1, 3, 5, 6, 7}); v != 6 || err != nil {
		t.Errorf("FindEven([1 3 5 6 7]) == (%d, %v), want (6, nil)", v, err)
	}
	if _, err := FindEven([]int{1, 3, 5, 7}); err == nil {
		t.Error("FindEven([1 3 5 7]) should return an error")
	}
}

func TestLookupAge(t *testing.T) {
	ages := map[string]int{"ada": 36, "grace": 85}
	if v, err := LookupAge(ages, "ada"); v != 36 || err != nil {
		t.Errorf("LookupAge(ages, \"ada\") == (%d, %v), want (36, nil)", v, err)
	}
	_, err := LookupAge(ages, "linus")
	if err == nil {
		t.Fatal("LookupAge(ages, \"linus\") should return an error")
	}
	if !strings.Contains(err.Error(), "linus") {
		t.Errorf("the error message %q should mention the missing name \"linus\"", err.Error())
	}
}
