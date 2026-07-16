package stringex

// Test code, don't change this!

import "testing"

func TestExercises(t *testing.T) {
	steps := []struct {
		name string
		run  func(*testing.T)
	}{
		{"01_Shout_worked_example", testShout},
		{"02_CountVowels", testCountVowels},
		{"03_Reverse", testReverse},
		{"04_IsPalindrome", testIsPalindrome},
		{"05_Initials", testInitials},
	}
	for _, step := range steps {
		if !t.Run(step.name, step.run) {
			return
		}
	}
}

func testShout(t *testing.T) {
	if got := Shout("go"); got != "GO!" {
		t.Errorf("Shout(\"go\")\n  got:  %q\n  want: %q", got, "GO!")
	}
}

func testCountVowels(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"hello", 2}, {"xyz", 0}, {"aeiou", 5}, {"", 0}, {"gopher gopher", 4},
	}
	for _, c := range cases {
		if got := CountVowels(c.in); got != c.want {
			t.Errorf("CountVowels(%q)\n  got:  %d\n  want: %d", c.in, got, c.want)
		}
	}
}

func testReverse(t *testing.T) {
	cases := []struct{ in, want string }{
		{"hello", "olleh"}, {"", ""}, {"a", "a"}, {"héllo", "olléh"},
	}
	for _, c := range cases {
		if got := Reverse(c.in); got != c.want {
			t.Errorf("Reverse(%q)\n  got:  %q\n  want: %q", c.in, got, c.want)
		}
	}
}

func testIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"racecar", true}, {"hello", false}, {"", true}, {"ab", false}, {"aba", true},
	}
	for _, c := range cases {
		if got := IsPalindrome(c.in); got != c.want {
			t.Errorf("IsPalindrome(%q)\n  got:  %v\n  want: %v", c.in, got, c.want)
		}
	}
}

func testInitials(t *testing.T) {
	cases := []struct{ in, want string }{
		{"grace brewster hopper", "GBH"},
		{"go", "G"},
		{"", ""},
		{"alan  turing", "AT"},
	}
	for _, c := range cases {
		if got := Initials(c.in); got != c.want {
			t.Errorf("Initials(%q)\n  got:  %q\n  want: %q", c.in, got, c.want)
		}
	}
}
