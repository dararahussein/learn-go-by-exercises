package main

// Test code, don't change this!

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// get sends a fake request to the mux and returns the response body.
func get(t *testing.T, path string) string {
	t.Helper()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	newMux().ServeHTTP(rec, req)
	return rec.Body.String()
}

func TestHealth(t *testing.T) {
	if got := get(t, "/health"); got != "ok" {
		t.Errorf("GET /health == %q, want \"ok\"", got)
	}
}

func TestGreetWithName(t *testing.T) {
	if got := get(t, "/greet?name=Ada"); got != "Hello, Ada!" {
		t.Errorf("GET /greet?name=Ada == %q, want \"Hello, Ada!\"", got)
	}
}

func TestGreetDefault(t *testing.T) {
	if got := get(t, "/greet"); got != "Hello, world!" {
		t.Errorf("GET /greet == %q, want \"Hello, world!\"", got)
	}
}
