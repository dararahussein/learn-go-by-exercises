package set01

// Test code, don't change this!

import "testing"

func TestShout(t *testing.T) {
	if got := Shout("go"); got != "GO!" {
		t.Errorf(`Shout("go") == %q, want "GO!"`, got)
	}
}

func TestCountVowels(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"hello", 2}, {"xyz", 0}, {"aeiou", 5}, {"", 0}, {"gopher gopher", 4},
	}
	for _, c := range cases {
		if got := CountVowels(c.in); got != c.want {
			t.Errorf("CountVowels(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestReverse(t *testing.T) {
	cases := []struct{ in, want string }{
		{"hello", "olleh"}, {"", ""}, {"a", "a"}, {"héllo", "olléh"},
	}
	for _, c := range cases {
		if got := Reverse(c.in); got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"racecar", true}, {"hello", false}, {"", true}, {"ab", false}, {"aba", true},
	}
	for _, c := range cases {
		if got := IsPalindrome(c.in); got != c.want {
			t.Errorf("IsPalindrome(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestInitials(t *testing.T) {
	cases := []struct{ in, want string }{
		{"grace brewster hopper", "GBH"},
		{"go", "G"},
		{"", ""},
		{"alan  turing", "AT"},
	}
	for _, c := range cases {
		if got := Initials(c.in); got != c.want {
			t.Errorf("Initials(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
