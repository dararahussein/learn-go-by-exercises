package debugging

import (
	"strings"
	"testing"
)

func TestCredentialsRoundTrip(t *testing.T) {
	data, err := CredentialsJSON(NewCredentials("ada", "analytical-engine"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"secret":"analytical-engine"`) {
		t.Errorf("JSON %s does not contain the secret field", data)
	}
	got, err := CredentialsFromJSON(data)
	if err != nil || got.SecretValue() != "analytical-engine" {
		t.Errorf("round trip = (%+v, %v); secret was lost", got, err)
	}
}
