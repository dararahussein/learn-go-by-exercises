package collections

import (
	"slices"
	"testing"
)

func TestFilterAndMap(t *testing.T) {
	even := Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })
	if !slices.Equal(even, []int{2, 4}) {
		t.Errorf("Filter = %v", even)
	}
	words := Map([]int{1, 2}, func(n int) string { return string(rune('a' + n - 1)) })
	if !slices.Equal(words, []string{"a", "b"}) {
		t.Errorf("Map = %v", words)
	}
}
