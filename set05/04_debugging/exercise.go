/*
DEBUGGING: encoding/json silently ignores unexported fields. Run `go vet` too:
its warning about a JSON tag on an unexported field is additional evidence.
*/
package debugging

import "encoding/json"

type Credentials struct {
	Username string `json:"username"`
	secret   string `json:"secret"` // BUG: lowercase fields are invisible to encoding/json
}

func NewCredentials(username, secret string) Credentials {
	return Credentials{Username: username, secret: secret}
}

func (c Credentials) SecretValue() string { return c.secret }

func CredentialsJSON(c Credentials) ([]byte, error) { return json.Marshal(c) }

func CredentialsFromJSON(data []byte) (Credentials, error) {
	var c Credentials
	err := json.Unmarshal(data, &c)
	return c, err
}

/*
Q1: Predict the exact JSON produced before running the test. Why is secret gone?
A1:
Q2: Missing JSON keys and explicit zero values both decode to Go zero values.
    When would a pointer field help distinguish them?
A2:
*/
