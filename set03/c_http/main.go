/*

A tiny HTTP server — the shape of every Go web service.

A handler is any function with this signature:
    func(w http.ResponseWriter, r *http.Request)
You WRITE the response into w; you READ the request from r. A ServeMux maps
URL paths to handlers. That's basically the whole model.

This is a runnable program:
    go run ./set03/c_http
then visit http://localhost:8080/health and http://localhost:8080/greet?name=Ada
Stop it with Ctrl+C.

Tasks:
- [ ] Implement greetHandler and register it in newMux.
- [ ] Run `go test ./set03/c_http/` — the tests hit the handlers in-memory
      (no real network needed), which is how Go web code is normally tested.

Expected time to complete: 45 minutes.

*/

package main

import (
	"fmt"
	"net/http"
)

// DONE FOR YOU: a handler that always replies "ok".
// fmt.Fprint writes to w, which implements io.Writer (the set02 idea again).
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

// greetHandler should reply "Hello, <name>!" where name comes from the URL
// query, e.g. /greet?name=Ada -> "Hello, Ada!". If name is missing or empty,
// use "world" -> "Hello, world!".
// Read a query parameter with: r.URL.Query().Get("name")
func greetHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: write "Hello, <name>!" to w, where name = r.URL.Query().Get("name")
	// (default "world" when empty). Until you do, this writes nothing and the
	// test fails.
}

// newMux wires paths to handlers. Tests call this directly, so keep the route
// registration here (not inside main).
func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	// TODO: register greetHandler at the path "/greet"
	return mux
}

func main() {
	fmt.Println("listening on http://localhost:8080")
	http.ListenAndServe(":8080", newMux())
}
