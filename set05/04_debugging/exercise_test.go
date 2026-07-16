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
		t.Fatalf("CredentialsJSON: got %s, want JSON containing %q", data, `"secret":"analytical-engine"`)
	}
	got, err := CredentialsFromJSON(data)
	if err != nil || got.SecretValue() != "analytical-engine" {
		t.Fatalf("credentials round trip: got (%+v, %v), want secret %q and nil error", got, err, "analytical-engine")
	}
}
