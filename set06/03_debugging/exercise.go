package debugging

import (
	"fmt"
	"io"
	"net/http"
)

// BUG: WriteHeader must happen before the response body is written.
// Predict the observed status, run the test, then fix the order.
func validationHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "invalid input")
	w.WriteHeader(http.StatusBadRequest)
}

// BUG: reading succeeded, but the response body was never closed. Closing is
// what permits connection reuse and releases resources owned by the transport.
func FetchMessage(client *http.Client, url string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
Q1: Who owns an http.Response.Body, and when must it be closed?
A1:
Q2: Why is *http.Client a parameter instead of calling http.Get directly?
A2:
*/
