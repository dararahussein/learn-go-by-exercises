package jsonex

import (
	"slices"
	"testing"
)

const peopleJSON = `[{"name":"Ada","age":36},{"name":"Grace","age":85},{"name":"Linus","age":20}]`

func TestExercises(t *testing.T) {
	if !t.Run("01_PersonJSON", testPersonJSON) {
		return
	}
	t.Run("02_PeopleQueries", testPeopleQueries)
}

func testPersonJSON(t *testing.T) {
	got, err := PersonToJSON(Person{Name: "Ada", Age: 36})
	if err != nil || got != `{"name":"Ada","age":36}` {
		t.Errorf("PersonToJSON\n  got:  (%s, %v)\n  want: (%s, nil)", got, err, `{"name":"Ada","age":36}`)
	}
	p, err := PersonFromJSON(`{"name":"Grace","age":85}`)
	if err != nil || p.Name != "Grace" || p.Age != 85 {
		t.Errorf("PersonFromJSON\n  got:  (%+v, %v)\n  want: ({Name:Grace Age:85}, nil)", p, err)
	}
	if _, err := PersonFromJSON(`not json`); err == nil {
		t.Error("PersonFromJSON(invalid) error\n  got:  nil\n  want: non-nil")
	}
}

func testPeopleQueries(t *testing.T) {
	if got, err := TotalAge(peopleJSON); err != nil || got != 141 {
		t.Errorf("TotalAge\n  got:  (%d, %v)\n  want: (141, nil)", got, err)
	}
	got, err := NamesOver(peopleJSON, 30)
	if err != nil || !slices.Equal(got, []string{"Ada", "Grace"}) {
		t.Errorf("NamesOver\n  got:  (%v, %v)\n  want: ([Ada Grace], nil)", got, err)
	}
}
