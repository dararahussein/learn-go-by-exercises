package application

import (
	"slices"
	"testing"
)

func TestUniqueWords(t *testing.T) {
	got := UniqueWords("Go gopher go  LEARN")
	want := []string{"go", "gopher", "learn"}
	if !slices.Equal(got, want) {
		t.Errorf("UniqueWords\n  got:  %v\n  want: %v", got, want)
	}
}
