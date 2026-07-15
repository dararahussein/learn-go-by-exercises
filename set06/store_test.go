package main

// Test code, don't change this!

import (
	"path/filepath"
	"testing"
)

func TestAddAssignsIDs(t *testing.T) {
	s := &Store{}
	a := s.Add("write go")
	b := s.Add("ship it")
	if a.ID != 1 || b.ID != 2 {
		t.Errorf("Add assigned IDs %d and %d, want 1 and 2", a.ID, b.ID)
	}
	if len(s.Tasks) != 2 {
		t.Fatalf("store has %d tasks, want 2", len(s.Tasks))
	}
	if s.Tasks[0].Text != "write go" || s.Tasks[1].Text != "ship it" {
		t.Errorf("stored texts = %q, %q", s.Tasks[0].Text, s.Tasks[1].Text)
	}
}

func TestPendingAndComplete(t *testing.T) {
	s := &Store{}
	s.Add("a")
	s.Add("b")
	if got := s.Pending(); len(got) != 2 {
		t.Fatalf("Pending == %d tasks, want 2", len(got))
	}
	if err := s.Complete(1); err != nil {
		t.Fatalf("Complete(1) unexpected error: %v", err)
	}
	pending := s.Pending()
	if len(pending) != 1 || pending[0].Text != "b" {
		t.Errorf("after completing task 1, Pending == %v, want just [b]", pending)
	}
}

func TestCompleteUnknownID(t *testing.T) {
	s := &Store{}
	s.Add("only task")
	if err := s.Complete(99); err == nil {
		t.Error("Complete(99) should return an error for a missing id")
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	s := &Store{}
	s.Add("persist me")
	s.Add("and me")
	_ = s.Complete(1)
	if err := s.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}

	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(loaded.Tasks) != 2 {
		t.Fatalf("loaded %d tasks, want 2", len(loaded.Tasks))
	}
	if !loaded.Tasks[0].Done {
		t.Error("completed state did not survive Save/Load")
	}
	if loaded.Tasks[1].Text != "and me" {
		t.Errorf("loaded text = %q, want \"and me\"", loaded.Tasks[1].Text)
	}
}

func TestLoadMissingFileIsEmpty(t *testing.T) {
	loaded, err := Load(filepath.Join(t.TempDir(), "nope.json"))
	if err != nil {
		t.Fatalf("Load of a missing file should NOT error, got: %v", err)
	}
	if loaded == nil || len(loaded.Tasks) != 0 {
		t.Errorf("Load of a missing file should give an empty store, got %v", loaded)
	}
}
