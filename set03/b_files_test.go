package set03

// Test code, don't change this!

import (
	"maps"
	"path/filepath"
	"slices"
	"testing"
)

// writeTemp creates a throwaway file for a test and returns its path.
// t.TempDir() is cleaned up automatically when the test finishes.
func writeTemp(t *testing.T, name, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	if err := WriteLines(path, []string{content}); err != nil {
		t.Fatalf("setup: %v", err)
	}
	return path
}

func TestReadLines(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "poem.txt")
	want := []string{"roses are red", "violets are blue"}
	if err := WriteLines(path, want); err != nil {
		t.Fatalf("setup: %v", err)
	}
	got, err := ReadLines(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !slices.Equal(got, want) {
		t.Errorf("ReadLines == %v, want %v", got, want)
	}
}

func TestCountLines(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "three.txt")
	if err := WriteLines(path, []string{"a", "b", "c"}); err != nil {
		t.Fatalf("setup: %v", err)
	}
	got, err := CountLines(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 3 {
		t.Errorf("CountLines == %d, want 3", got)
	}
}

func TestWordFrequency(t *testing.T) {
	path := writeTemp(t, "words.txt", "go go   gopher\nGo")
	got, err := WordFrequency(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := map[string]int{"go": 2, "gopher": 1, "Go": 1}
	if !maps.Equal(got, want) {
		t.Errorf("WordFrequency == %v, want %v", got, want)
	}
}

func TestReadLinesMissingFile(t *testing.T) {
	if _, err := ReadLines(filepath.Join(t.TempDir(), "does-not-exist.txt")); err == nil {
		t.Error("ReadLines on a missing file should return an error")
	}
}
