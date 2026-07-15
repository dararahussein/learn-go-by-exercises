package set06

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchPeople(t *testing.T) {
	server := httptest.NewServer(NewMux())
	t.Cleanup(server.Close)
	got, err := FetchPeople(server.Client(), server.URL+"/people")
	if err != nil || len(got) != 2 || got[0].Name != "Ada" {
		t.Errorf("FetchPeople = (%v, %v)", got, err)
	}
}

func TestFetchPeopleRejectsErrorStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "try later", http.StatusServiceUnavailable)
	}))
	t.Cleanup(server.Close)
	if _, err := FetchPeople(server.Client(), server.URL); err == nil {
		t.Error("FetchPeople should reject a 503 response")
	}
}
