package debugging

// Test code, don't change this!

import "testing"

func TestExercises(t *testing.T) {
	steps := []struct {
		name string
		run  func(*testing.T)
	}{
		{"01_LastIndexOf", testLastIndexOf},
		{"02_Contains", testContains},
		{"03_Average", testAverage},
	}
	for _, step := range steps {
		step.run(t)
	}
}

func testLastIndexOf(t *testing.T) {
	cases := []struct {
		xs     []int
		target int
		want   int
	}{
		{[]int{10, 20, 10}, 10, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{5, 5, 5}, 5, 2},
		{[]int{1, 2, 3}, 9, -1},
		{nil, 1, -1},
	}
	for _, c := range cases {
		if got := LastIndexOf(c.xs, c.target); got != c.want {
			t.Fatalf("LastIndexOf(%v, %d): got %d, want %d", c.xs, c.target, got, c.want)
		}
	}
}

func testContains(t *testing.T) {
	cases := []struct {
		xs     []int
		target int
		want   bool
	}{
		{[]int{1, 2, 3}, 3, true},
		{[]int{1, 2, 3}, 1, true},
		{[]int{1, 2, 3}, 9, false},
		{nil, 1, false},
	}
	for _, c := range cases {
		if got := Contains(c.xs, c.target); got != c.want {
			t.Fatalf("Contains(%v, %d): got %v, want %v", c.xs, c.target, got, c.want)
		}
	}
}

func testAverage(t *testing.T) {
	cases := []struct {
		xs   []int
		want float64
	}{
		{[]int{1, 2}, 1.5},
		{[]int{2, 2, 2}, 2},
		{[]int{1, 2, 3, 4}, 2.5},
		{nil, 0},
	}
	for _, c := range cases {
		if got := Average(c.xs); got != c.want {
			t.Fatalf("Average(%v): got %v, want %v", c.xs, got, c.want)
		}
	}
}
