package api

import (
	"encoding/json"
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
	var got []task.Task
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("decode GET /tasks: %v", err)
	}
	if len(got) != 1 || got[0].ID != 1 || got[0].Text != "learn Go" || got[0].Done {
		t.Errorf("GET tasks = %+v; want one pending learn Go task", got)
	}
}

func TestPostTasks(t *testing.T) {
	// TODO: replace this failure with table-driven cases BEFORE implementing POST:
	// valid JSON -> 201, malformed JSON -> 400, empty text -> 400.
	t.Fatal("write POST /tasks tests first")
}
