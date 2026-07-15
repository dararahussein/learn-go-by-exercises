package set05

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigPreservesCauses(t *testing.T) {
	_, err := LoadConfig(filepath.Join(t.TempDir(), "missing.json"))
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("missing config error %v should preserve os.ErrNotExist", err)
	}
	path := filepath.Join(t.TempDir(), "bad.json")
	if err := os.WriteFile(path, []byte(`{"address":`), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err = LoadConfig(path)
	var syntaxErr *json.SyntaxError
	if !errors.As(err, &syntaxErr) {
		t.Errorf("decode error %v should preserve *json.SyntaxError", err)
	}
}
