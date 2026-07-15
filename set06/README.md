# Set 06 - HTTP servers and clients

Handlers are ordinary functions and can be tested without opening a real port.
Clients own response bodies and should accept dependencies that make timeouts
and tests controllable.

Warm-up from memory: write an error wrap that adds an operation name while
preserving the cause for `errors.Is`.

Implement the handlers first, then use `httptest.NewServer` through
`FetchPeople`. The two debugging tasks cover response-header ordering and a
resource leak. Run the program manually with `go run ./set06/cmd/server` only
after its targeted tests are green.

Optional stretch: add logging middleware that records method, path, status, and
duration without changing the handlers.

Expected time: 4-6 hours.
