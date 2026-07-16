/*
HTTP handlers are ordinary functions. Tests can call them in memory with
httptest, so a network port is never required for the feedback loop.
*/
package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people = []Person{{Name: "Ada", Age: 36}, {Name: "Grace", Age: 85}}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "ok")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: write "Hello, <name>!"; default an empty name to "world"
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: GET /people returns people as JSON with Content-Type application/json.
	// If ?name=X is supplied, return that one person or http.StatusNotFound.
}

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	// TODO: register /greet and /people
	return mux
}

// Keep json in the starter import list while the handler is a stub.
var _ = json.NewEncoder
