package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	serverexercise "learn-go-by-exercises/set06/01_server"
)

func TestExercises(t *testing.T) {
	testFetchPeople(t)
	testFetchPeopleRejectsErrorStatus(t)
}

func testFetchPeople(t *testing.T) {
	server := httptest.NewServer(serverexercise.NewMux())
	t.Cleanup(server.Close)
	got, err := FetchPeople(server.Client(), server.URL+"/people")
	if err != nil || len(got) != 2 || got[0].Name != "Ada" {
		t.Fatalf("FetchPeople: got (%v, %v), want (2 people starting with Ada, nil)", got, err)
	}
}

func testFetchPeopleRejectsErrorStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "try later", http.StatusServiceUnavailable)
	}))
	t.Cleanup(server.Close)
	if _, err := FetchPeople(server.Client(), server.URL); err == nil {
		t.Fatal("FetchPeople(503) error: got nil, want non-nil")
	}
}
