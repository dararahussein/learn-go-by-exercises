package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

var ErrNotFound = errors.New("task not found")

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type Store struct {
	mu    sync.RWMutex
	Tasks []Task `json:"tasks"`
}

// Save is done for you. Notice that even JSON encoding counts as a read.
func (s *Store) Save(path string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("encode task store: %w", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write task store %q: %w", path, err)
	}
	return nil
}

// Load treats a missing file as a new empty store and wraps other failures.
func Load(path string) (*Store, error) {
	return &Store{}, nil // TODO
}

func (s *Store) Add(text string) Task {
	return Task{} // TODO: highest existing ID + 1; protect the mutation
}

// BUG: this returns success but modifies a range copy. Diagnose and fix it.
func (s *Store) Complete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, item := range s.Tasks {
		if item.ID == id {
			item.Done = true
			return nil
		}
	}
	return fmt.Errorf("complete task %d: %w", id, ErrNotFound)
}

func (s *Store) Pending() []Task {
	return nil // TODO: protect read; return a new filtered slice
}
