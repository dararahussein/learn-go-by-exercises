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
		t.Errorf("CredentialsJSON\n  got:  %s\n  want: JSON containing %q", data, `"secret":"analytical-engine"`)
	}
	got, err := CredentialsFromJSON(data)
	if err != nil || got.SecretValue() != "analytical-engine" {
		t.Errorf("credentials round trip\n  got:  (%+v, %v)\n  want: secret %q and nil error", got, err, "analytical-engine")
	}
}
