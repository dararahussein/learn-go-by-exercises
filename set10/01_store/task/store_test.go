package task

import (
	"errors"
	"path/filepath"
	"sync"
	"testing"
)

func TestExercises(t *testing.T) {
	if !t.Run("01_AddCompletePending", testAddCompletePending) {
		return
	}
	if !t.Run("02_SaveLoadRoundTrip", testSaveLoadRoundTrip) {
		return
	}
	t.Run("03_ConcurrentAdd", testConcurrentAdd)
}

func testAddCompletePending(t *testing.T) {
	s := &Store{}
	a := s.Add("write Go")
	b := s.Add("ship it")
	if a.ID != 1 || b.ID != 2 {
		t.Fatalf("IDs = %d, %d; want 1, 2", a.ID, b.ID)
	}
	if err := s.Complete(1); err != nil {
		t.Fatal(err)
	}
	pending := s.Pending()
	if len(pending) != 1 || pending[0].Text != "ship it" {
		t.Errorf("Pending = %+v; want only ship it", pending)
	}
	if err := s.Complete(99); !errors.Is(err, ErrNotFound) {
		t.Errorf("Complete(99) error %v should wrap ErrNotFound", err)
	}
}

func testSaveLoadRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	s := &Store{}
	s.Add("persist me")
	if err := s.Save(path); err != nil {
		t.Fatal(err)
	}
	got, err := Load(path)
	if err != nil || len(got.Tasks) != 1 || got.Tasks[0].Text != "persist me" {
		t.Errorf("Load = (%+v, %v)", got, err)
	}
	missing, err := Load(filepath.Join(t.TempDir(), "missing.json"))
	if err != nil || missing == nil || len(missing.Tasks) != 0 {
		t.Errorf("Load(missing) = (%+v, %v); want empty store", missing, err)
	}
}

func testConcurrentAdd(t *testing.T) {
	s := &Store{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); s.Add("parallel") }()
	}
	wg.Wait()
	if got := len(s.Pending()); got != 100 {
		t.Errorf("Pending count = %d; want 100", got)
	}
}
