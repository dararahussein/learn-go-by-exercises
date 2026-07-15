package set01

import (
	"slices"
	"testing"
)

func TestUniqueWords(t *testing.T) {
	got := UniqueWords("Go gopher go  LEARN")
	want := []string{"go", "gopher", "learn"}
	if !slices.Equal(got, want) {
		t.Errorf("UniqueWords = %v; want %v", got, want)
	}
}
