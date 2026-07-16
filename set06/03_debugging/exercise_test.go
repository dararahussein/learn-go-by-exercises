package debugging

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type trackingBody struct {
	io.Reader
	closed bool
}

func (b *trackingBody) Close() error { b.closed = true; return nil }

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestExercises(t *testing.T) {
	testValidationStatus(t)
	testFetchMessageClosesBody(t)
}

func testValidationStatus(t *testing.T) {
	rec := httptest.NewRecorder()
	validationHandler(rec, httptest.NewRequest(http.MethodPost, "/", nil))
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status: got %d, want 400", rec.Code)
	}
}

func testFetchMessageClosesBody(t *testing.T) {
	body := &trackingBody{Reader: strings.NewReader("hello")}
	client := &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusOK, Body: body, Header: make(http.Header)}, nil
	})}
	got, err := FetchMessage(client, "http://example.test")
	if err != nil || got != "hello" {
		t.Fatalf("FetchMessage: got (%q, %v), want (%q, nil)", got, err, "hello")
	}
	if !body.closed {
		t.Fatal("response body closed: got false, want true")
	}
}
