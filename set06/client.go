package set06

import (
	"fmt"
	"net/http"
)

// FetchPeople must GET url, reject non-2xx statuses, close the response body,
// and decode the JSON array. Accepting a client makes it testable and lets a
// production caller configure timeouts.
func FetchPeople(client *http.Client, url string) ([]Person, error) {
	return nil, fmt.Errorf("TODO: FetchPeople %s", url)
}
