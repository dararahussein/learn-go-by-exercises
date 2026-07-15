package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"learn-go-by-exercises/set10/task"
)

func TestGetTasks(t *testing.T) {
	store := &task.Store{}
	store.Add("learn Go")
	rec := httptest.NewRecorder()
	NewHandler(store).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/tasks", nil))
	if rec.Code != http.StatusOK || rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("GET status/type = %d/%q", rec.Code, rec.Header().Get("Content-Type"))
	}
	if body := rec.Body.String(); body == "" || body == "null\n" {
		t.Errorf("GET body = %q; want tasks JSON", body)
	}
}

func TestPostTasks(t *testing.T) {
	// TODO: replace this failure with table-driven cases BEFORE implementing POST:
	// valid JSON -> 201, malformed JSON -> 400, empty text -> 400.
	t.Fatal("write POST /tasks tests first")
}
