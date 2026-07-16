package debugging

import "testing"

func TestParsePort(t *testing.T) {
	if got, err := ParsePort("8080"); err != nil || got != 8080 {
		t.Fatalf("ParsePort(8080): got (%d, %v), want (8080, nil)", got, err)
	}
	if _, err := ParsePort("not-a-port"); err == nil {
		t.Fatal("ParsePort(non-number) error: got nil, want non-nil")
	}
}
