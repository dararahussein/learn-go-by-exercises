/* JSON maps exported struct fields to object keys through struct tags. */
package set05

import "encoding/json"

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func PersonToJSON(p Person) (string, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func PersonFromJSON(data string) (Person, error) {
	return Person{}, nil // TODO: unmarshal into a Person
}

func TotalAge(data string) (int, error) {
	return 0, nil // TODO: decode []Person and sum ages
}

func NamesOver(data string, minAge int) ([]string, error) {
	return nil, nil // TODO: decode and filter in input order
}
