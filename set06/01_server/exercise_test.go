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
	if !t.Run("01_HealthAndGreet", testHealthAndGreet) {
		return
	}
	t.Run("02_PeopleJSON", testPeopleHandler)
}

func testHealthAndGreet(t *testing.T) {
	if got := request(t, "/health").Body.String(); got != "ok" {
		t.Errorf("GET /health body\n  got:  %q\n  want: %q", got, "ok")
	}
	for path, want := range map[string]string{
		"/greet?name=Ada": "Hello, Ada!",
		"/greet":          "Hello, world!",
	} {
		if got := request(t, path).Body.String(); got != want {
			t.Errorf("GET %s body\n  got:  %q\n  want: %q", path, got, want)
		}
	}
}

func testPeopleHandler(t *testing.T) {
	rec := request(t, "/people")
	if rec.Code != http.StatusOK || rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("GET /people status/content-type\n  got:  %d/%q\n  want: 200/%q", rec.Code, rec.Header().Get("Content-Type"), "application/json")
	}
	var got []Person
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil || len(got) != 2 {
		t.Errorf("GET /people decoded value\n  got:  (%v, %v)\n  want: (2 people, nil)", got, err)
	}
	if got := request(t, "/people?name=missing").Code; got != http.StatusNotFound {
		t.Errorf("missing person status\n  got:  %d\n  want: 404", got)
	}
}
