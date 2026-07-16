package jsonex

import (
	"slices"
	"testing"
)

const peopleJSON = `[{"name":"Ada","age":36},{"name":"Grace","age":85},{"name":"Linus","age":20}]`

func TestExercises(t *testing.T) {
	testPersonJSON(t)
	testPeopleQueries(t)
}

func testPersonJSON(t *testing.T) {
	got, err := PersonToJSON(Person{Name: "Ada", Age: 36})
	if err != nil || got != `{"name":"Ada","age":36}` {
		t.Fatalf("PersonToJSON: got (%s, %v), want (%s, nil)", got, err, `{"name":"Ada","age":36}`)
	}
	p, err := PersonFromJSON(`{"name":"Grace","age":85}`)
	if err != nil || p.Name != "Grace" || p.Age != 85 {
		t.Fatalf("PersonFromJSON: got (%+v, %v), want ({Name:Grace Age:85}, nil)", p, err)
	}
	if _, err := PersonFromJSON(`not json`); err == nil {
		t.Fatal("PersonFromJSON(invalid) error: got nil, want non-nil")
	}
}

func testPeopleQueries(t *testing.T) {
	if got, err := TotalAge(peopleJSON); err != nil || got != 141 {
		t.Fatalf("TotalAge: got (%d, %v), want (141, nil)", got, err)
	}
	got, err := NamesOver(peopleJSON, 30)
	if err != nil || !slices.Equal(got, []string{"Ada", "Grace"}) {
		t.Fatalf("NamesOver: got (%v, %v), want ([Ada Grace], nil)", got, err)
	}
}
