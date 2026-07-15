package set06

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

func TestHealthAndGreet(t *testing.T) {
	if got := request(t, "/health").Body.String(); got != "ok" {
		t.Errorf("GET /health body = %q; want ok", got)
	}
	for path, want := range map[string]string{
		"/greet?name=Ada": "Hello, Ada!",
		"/greet":          "Hello, world!",
	} {
		if got := request(t, path).Body.String(); got != want {
			t.Errorf("GET %s body = %q; want %q", path, got, want)
		}
	}
}

func TestPeopleHandler(t *testing.T) {
	rec := request(t, "/people")
	if rec.Code != http.StatusOK || rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("status/content-type = %d/%q", rec.Code, rec.Header().Get("Content-Type"))
	}
	var got []Person
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil || len(got) != 2 {
		t.Errorf("decode people = (%v, %v)", got, err)
	}
	if got := request(t, "/people?name=missing").Code; got != http.StatusNotFound {
		t.Errorf("missing person status = %d; want 404", got)
	}
}
