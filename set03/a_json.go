/*

JSON — how your Go program talks to the outside world (APIs, config, files).

The whole model: a struct's fields map to JSON keys, controlled by `json:"..."`
tags. encoding/json converts both directions:
  json.Marshal(value)     -> []byte      (Go -> JSON)
  json.Unmarshal(data, &v) -> fills &v    (JSON -> Go)

Note the & : Unmarshal writes INTO your variable, so it needs a pointer.

Tasks:
- [ ] Replace each panic("TODO...") with a real implementation.
- [ ] Run `go test ./set03/` after every change.

Expected time to complete: 45 minutes.

*/

package set03

import "encoding/json"

// The tags rename fields for JSON. `omitempty` drops the field when it's the
// zero value — so a Person with no email marshals without an "email" key.
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

// DONE FOR YOU: Go value -> JSON string.
func PersonToJSON(p Person) (string, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// PersonFromJSON parses one Person out of a JSON string. Declare a Person,
// json.Unmarshal([]byte(data), &p), propagate any error, return the person.
func PersonFromJSON(data string) (Person, error) {
	return Person{}, nil // TODO: implement PersonFromJSON
}

// TotalAge parses a JSON ARRAY of people and sums their ages.
// Unmarshalling into a []Person works exactly the same way — the target just
// happens to be a slice. Input looks like: [{"name":"a","age":1}, ...]
func TotalAge(data string) (int, error) {
	return 0, nil // TODO: implement TotalAge
}

// NamesOver parses a JSON array of people and returns the names of everyone
// whose age is >= minAge, in input order. Combines set01 slice-building with
// JSON parsing — this is what real request handlers look like.
func NamesOver(data string, minAge int) ([]string, error) {
	return nil, nil // TODO: implement NamesOver
}
