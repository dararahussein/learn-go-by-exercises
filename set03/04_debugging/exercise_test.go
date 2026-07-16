package debugging

import "testing"

func TestParsePort(t *testing.T) {
	if got, err := ParsePort("8080"); err != nil || got != 8080 {
		t.Errorf("ParsePort(8080) = (%d, %v); want (8080, nil)", got, err)
	}
	if _, err := ParsePort("not-a-port"); err == nil {
		t.Error("ParsePort(non-number) should return an error")
	}
}
