package files

import (
	"maps"
	"path/filepath"
	"slices"
	"testing"
)

func TestExercises(t *testing.T) {
	steps := []struct {
		name string
		run  func(*testing.T)
	}{
		{"01_ReadLines", testReadLines},
		{"02_WordFrequency", testWordFrequency},
		{"03_CountLines", testCountLines},
		{"04_MissingFile", testReadLinesMissing},
	}
	for _, step := range steps {
		if !t.Run(step.name, step.run) {
			return
		}
	}
}

func testReadLines(t *testing.T) {
	path := filepath.Join(t.TempDir(), "poem.txt")
	want := []string{"roses are red", "violets are blue"}
	if err := WriteLines(path, want); err != nil {
		t.Fatal(err)
	}
	got, err := ReadLines(path)
	if err != nil || !slices.Equal(got, want) {
		t.Errorf("ReadLines\n  got:  (%v, %v)\n  want: (%v, nil)", got, err, want)
	}
}

func testWordFrequency(t *testing.T) {
	path := filepath.Join(t.TempDir(), "words.txt")
	if err := WriteLines(path, []string{"go go", "gopher Go"}); err != nil {
		t.Fatal(err)
	}
	got, err := WordFrequency(path)
	want := map[string]int{"go": 2, "gopher": 1, "Go": 1}
	if err != nil || !maps.Equal(got, want) {
		t.Errorf("WordFrequency\n  got:  (%v, %v)\n  want: %v", got, err, want)
	}
}

func testCountLines(t *testing.T) {
	path := filepath.Join(t.TempDir(), "three.txt")
	if err := WriteLines(path, []string{"one", "two", "three"}); err != nil {
		t.Fatal(err)
	}
	if got, err := CountLines(path); err != nil || got != 3 {
		t.Errorf("CountLines\n  got:  (%d, %v)\n  want: (3, nil)", got, err)
	}
}

func testReadLinesMissing(t *testing.T) {
	if _, err := ReadLines(filepath.Join(t.TempDir(), "missing")); err == nil {
		t.Error("ReadLines(missing) error\n  got:  nil\n  want: non-nil")
	}
}

// Add one test yourself: verify the behavior of ReadLines on an empty file.
