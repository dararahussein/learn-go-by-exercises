/*

CAPSTONE — a task-tracker CLI (like a tiny `todo`).

This set has no hand-holding header per function: you've met every idea already
(structs, methods, errors, slices, JSON, files). Now you assemble them.

Read set06/README.md first for the spec and how to run it.

store.go holds the LOGIC (and is unit-tested in store_test.go).
main.go wires that logic to the command line.

Work order:
  1. Make `go test ./set06/` green by implementing the methods below.
  2. Then implement the commands in main.go and drive it with `go run ./set06 ...`.

*/

package main

import (
	"encoding/json"
	"os"
)

// A single task. JSON tags control how it's saved to disk.
type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// Store is the whole task list. It marshals straight to/from the JSON file.
type Store struct {
	Tasks []Task `json:"tasks"`
}

// DONE FOR YOU: write the store to disk as pretty JSON.
// MarshalIndent is like Marshal but human-readable — nice for a file you might
// open in an editor.
func (s *Store) Save(path string) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// Load reads a Store from disk.
// IMPORTANT real-world behavior: if the file doesn't exist yet (first run),
// return an empty &Store{} and NO error — a missing task list isn't a failure.
// Hint: os.ReadFile returns an error you can test with
//   errors.Is(err, os.ErrNotExist)
// (you'll need to import "errors"). Otherwise json.Unmarshal the bytes into a Store.
func Load(path string) (*Store, error) {
	return &Store{}, nil // TODO: implement Load (missing file -> empty store, no error)
}

// Add appends a new task whose ID is one higher than the highest existing ID
// (so the first task gets ID 1). It returns the task it created.
func (s *Store) Add(text string) Task {
	return Task{} // TODO: implement Add
}

// Complete marks the task with the given id as Done.
// Return an error (mentioning the id) if no such task exists.
// Careful: to mutate a task inside s.Tasks, index into the slice
// (s.Tasks[i].Done = true) rather than copying it in a range loop.
func (s *Store) Complete(id int) error {
	return nil // TODO: implement Complete (error if id not found)
}

// Pending returns just the tasks that aren't done yet, in order.
func (s *Store) Pending() []Task {
	return nil // TODO: implement Pending
}
