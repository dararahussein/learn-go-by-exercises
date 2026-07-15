package set03

// Test code, don't change this!

import (
	"slices"
	"testing"
)

func TestPersonToJSON(t *testing.T) {
	got, err := PersonToJSON(Person{Name: "Ada", Age: 36})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := `{"name":"Ada","age":36}` // email omitted by omitempty
	if got != want {
		t.Errorf("PersonToJSON == %s, want %s", got, want)
	}
}

func TestPersonFromJSON(t *testing.T) {
	p, err := PersonFromJSON(`{"name":"Grace","age":85,"email":"g@navy.mil"}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := Person{Name: "Grace", Age: 85, Email: "g@navy.mil"}
	if p != want {
		t.Errorf("PersonFromJSON == %+v, want %+v", p, want)
	}
	if _, err := PersonFromJSON(`not json`); err == nil {
		t.Error("PersonFromJSON(\"not json\") should return an error")
	}
}

const people = `[{"name":"Ada","age":36},{"name":"Grace","age":85},{"name":"Linus","age":20}]`

func TestTotalAge(t *testing.T) {
	got, err := TotalAge(people)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 141 {
		t.Errorf("TotalAge == %d, want 141", got)
	}
}

func TestNamesOver(t *testing.T) {
	got, err := NamesOver(people, 30)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"Ada", "Grace"}
	if !slices.Equal(got, want) {
		t.Errorf("NamesOver(people, 30) == %v, want %v", got, want)
	}
}
