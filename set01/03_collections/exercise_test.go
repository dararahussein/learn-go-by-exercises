package collections

// Test code, don't change this!

import (
	"maps"
	"slices"
	"testing"
)

func TestExercises(t *testing.T) {
	steps := []struct {
		name string
		run  func(*testing.T)
	}{
		{"01_Doubled_worked_example", testDoubled},
		{"02_ArrayTotal", testArrayTotal},
		{"03_Sum", testSum},
		{"04_Evens", testEvens},
		{"05_IndexOf", testIndexOf},
		{"06_WordCount", testWordCount},
		{"07_Lookup", testLookup},
		{"08_SortedKeys", testSortedKeys},
	}
	for _, step := range steps {
		if !t.Run(step.name, step.run) {
			return
		}
	}
}

func testDoubled(t *testing.T) {
	if got := Doubled([]int{1, 2, 3}); !slices.Equal(got, []int{2, 4, 6}) {
		t.Errorf("Doubled([1 2 3])\n  got:  %v\n  want: [2 4 6]", got)
	}
}

func testArrayTotal(t *testing.T) {
	if got := ArrayTotal([3]int{10, 20, 30}); got != 60 {
		t.Errorf("ArrayTotal\n  got:  %d\n  want: 60", got)
	}
}

func testSum(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3}, 6}, {nil, 0}, {[]int{-1, 1}, 0}, {[]int{42}, 42},
	}
	for _, c := range cases {
		if got := Sum(c.in); got != c.want {
			t.Errorf("Sum(%v)\n  got:  %d\n  want: %d", c.in, got, c.want)
		}
	}
}

func testEvens(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{[]int{1, 3, 5}, nil},
		{nil, nil},
		{[]int{2}, []int{2}},
	}
	for _, c := range cases {
		if got := Evens(c.in); !slices.Equal(got, c.want) {
			t.Errorf("Evens(%v)\n  got:  %v\n  want: %v", c.in, got, c.want)
		}
	}
}

func testIndexOf(t *testing.T) {
	cases := []struct {
		xs     []int
		target int
		want   int
	}{
		{[]int{10, 20, 30}, 20, 1},
		{[]int{10, 20, 30}, 99, -1},
		{[]int{5, 5, 5}, 5, 0},
		{nil, 1, -1},
	}
	for _, c := range cases {
		if got := IndexOf(c.xs, c.target); got != c.want {
			t.Errorf("IndexOf(%v, %d)\n  got:  %d\n  want: %d", c.xs, c.target, got, c.want)
		}
	}
}

func testWordCount(t *testing.T) {
	got := WordCount([]string{"go", "is", "go"})
	want := map[string]int{"go": 2, "is": 1}
	if !maps.Equal(got, want) {
		t.Errorf("WordCount\n  got:  %v\n  want: %v", got, want)
	}
	if got := WordCount(nil); len(got) != 0 {
		t.Errorf("WordCount(nil)\n  got:  %v\n  want: empty", got)
	}
}

func testLookup(t *testing.T) {
	m := map[string]int{"hi": 1, "hello": 2}
	if v, ok := Lookup(m, "hello"); v != 2 || !ok {
		t.Errorf("Lookup(m, \"hello\")\n  got:  (%d, %v)\n  want: (2, true)", v, ok)
	}
	if v, ok := Lookup(m, "nope"); v != 0 || ok {
		t.Errorf("Lookup(m, \"nope\")\n  got:  (%d, %v)\n  want: (0, false)", v, ok)
	}
}

func testSortedKeys(t *testing.T) {
	got := SortedKeys(map[string]int{"cherry": 1, "apple": 2, "banana": 3})
	want := []string{"apple", "banana", "cherry"}
	if !slices.Equal(got, want) {
		t.Errorf("SortedKeys\n  got:  %v\n  want: %v", got, want)
	}
}
