package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func request(t *testing.T, path string) *httptest.ResponseRecorder {
	t.Helper()
	rec := httptest.NewRecorder()
	NewMux().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	return rec
}

func TestExercises(t *testing.T) {
	testHealthAndGreet(t)
	testPeopleHandler(t)
}

func testHealthAndGreet(t *testing.T) {
	if got := request(t, "/health").Body.String(); got != "ok" {
		t.Fatalf("GET /health body: got %q, want %q", got, "ok")
	}
	for path, want := range map[string]string{
		"/greet?name=Ada": "Hello, Ada!",
		"/greet":          "Hello, world!",
	} {
		if got := request(t, path).Body.String(); got != want {
			t.Fatalf("GET %s body: got %q, want %q", path, got, want)
		}
	}
}

func testPeopleHandler(t *testing.T) {
	rec := request(t, "/people")
	if rec.Code != http.StatusOK || rec.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("GET /people status/content-type: got %d/%q, want 200/%q", rec.Code, rec.Header().Get("Content-Type"), "application/json")
	}
	var got []Person
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil || len(got) != 2 {
		t.Fatalf("GET /people decoded value: got (%v, %v), want (2 people, nil)", got, err)
	}
	if got := request(t, "/people?name=missing").Code; got != http.StatusNotFound {
		t.Fatalf("missing person status: got %d, want 404", got)
	}
}
