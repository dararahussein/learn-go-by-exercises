package config

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
		t.Fatalf("missing config error: got %v, want error wrapping os.ErrNotExist", err)
	}
	path := filepath.Join(t.TempDir(), "bad.json")
	if err := os.WriteFile(path, []byte(`{"address":`), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err = LoadConfig(path)
	var syntaxErr *json.SyntaxError
	if !errors.As(err, &syntaxErr) {
		t.Fatalf("decode config error: got %v, want error wrapping *json.SyntaxError", err)
	}
}
