package client

import (
	"fmt"
	"net/http"

	server "learn-go-by-exercises/set06/01_server"
)

type Person = server.Person

// FetchPeople must GET url, reject non-2xx statuses, close the response body,
// and decode the JSON array. Accepting a client makes it testable and lets a
// production caller configure timeouts.
func FetchPeople(client *http.Client, url string) ([]Person, error) {
	return nil, fmt.Errorf("TODO: FetchPeople %s", url)
}
