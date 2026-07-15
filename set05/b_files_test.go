package set05

import (
	"maps"
	"path/filepath"
	"slices"
	"testing"
)

func TestReadLines(t *testing.T) {
	path := filepath.Join(t.TempDir(), "poem.txt")
	want := []string{"roses are red", "violets are blue"}
	if err := WriteLines(path, want); err != nil {
		t.Fatal(err)
	}
	got, err := ReadLines(path)
	if err != nil || !slices.Equal(got, want) {
		t.Errorf("ReadLines = (%v, %v); want (%v, nil)", got, err, want)
	}
}

func TestWordFrequency(t *testing.T) {
	path := filepath.Join(t.TempDir(), "words.txt")
	if err := WriteLines(path, []string{"go go", "gopher Go"}); err != nil {
		t.Fatal(err)
	}
	got, err := WordFrequency(path)
	want := map[string]int{"go": 2, "gopher": 1, "Go": 1}
	if err != nil || !maps.Equal(got, want) {
		t.Errorf("WordFrequency = (%v, %v); want %v", got, err, want)
	}
}

func TestCountLines(t *testing.T) {
	path := filepath.Join(t.TempDir(), "three.txt")
	if err := WriteLines(path, []string{"one", "two", "three"}); err != nil {
		t.Fatal(err)
	}
	if got, err := CountLines(path); err != nil || got != 3 {
		t.Errorf("CountLines = (%d, %v); want (3, nil)", got, err)
	}
}

func TestReadLinesMissing(t *testing.T) {
	if _, err := ReadLines(filepath.Join(t.TempDir(), "missing")); err == nil {
		t.Error("reading a missing file should return an error")
	}
}

// Add one test yourself: verify the behavior of ReadLines on an empty file.
