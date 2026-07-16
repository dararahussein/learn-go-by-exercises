package functions

// Test code, don't change this!

import "testing"

func TestExercises(t *testing.T) {
	steps := []struct {
		name string
		run  func(*testing.T)
	}{
		{"01_Double_worked_example", testDouble},
		{"02_Abs", testAbs},
		{"03_Max", testMax},
		{"04_Clamp", testClamp},
		{"05_DivMod", testDivMod},
		{"06_SumTo", testSumTo},
		{"07_FizzBuzz", testFizzBuzz},
	}
	for _, step := range steps {
		if !t.Run(step.name, step.run) {
			return
		}
	}
}

func testDouble(t *testing.T) {
	if got := Double(21); got != 42 {
		t.Errorf("Double(21) == %d, want 42", got)
	}
}

func testAbs(t *testing.T) {
	cases := []struct{ in, want int }{{-3, 3}, {3, 3}, {0, 0}, {-100, 100}}
	for _, c := range cases {
		if got := Abs(c.in); got != c.want {
			t.Errorf("Abs(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func testMax(t *testing.T) {
	cases := []struct{ x, y, want int }{{1, 2, 2}, {2, 1, 2}, {-5, -3, -3}, {7, 7, 7}}
	for _, c := range cases {
		if got := Max(c.x, c.y); got != c.want {
			t.Errorf("Max(%d, %d) == %d, want %d", c.x, c.y, got, c.want)
		}
	}
}

func testClamp(t *testing.T) {
	cases := []struct{ x, lo, hi, want int }{
		{5, 0, 10, 5}, {-5, 0, 10, 0}, {15, 0, 10, 10}, {0, 0, 10, 0}, {10, 0, 10, 10},
	}
	for _, c := range cases {
		if got := Clamp(c.x, c.lo, c.hi); got != c.want {
			t.Errorf("Clamp(%d, %d, %d) == %d, want %d", c.x, c.lo, c.hi, got, c.want)
		}
	}
}

func testDivMod(t *testing.T) {
	q, r := DivMod(17, 5)
	if q != 3 || r != 2 {
		t.Errorf("DivMod(17, 5) == (%d, %d), want (3, 2)", q, r)
	}
	q, r = DivMod(10, 2)
	if q != 5 || r != 0 {
		t.Errorf("DivMod(10, 2) == (%d, %d), want (5, 0)", q, r)
	}
}

func testSumTo(t *testing.T) {
	cases := []struct{ in, want int }{{1, 1}, {5, 15}, {100, 5050}, {0, 0}}
	for _, c := range cases {
		if got := SumTo(c.in); got != c.want {
			t.Errorf("SumTo(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}

func testFizzBuzz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, "1"}, {3, "Fizz"}, {5, "Buzz"}, {15, "FizzBuzz"}, {7, "7"}, {30, "FizzBuzz"}, {9, "Fizz"}, {10, "Buzz"},
	}
	for _, c := range cases {
		if got := FizzBuzz(c.in); got != c.want {
			t.Errorf("FizzBuzz(%d) == %q, want %q", c.in, got, c.want)
		}
	}
}
