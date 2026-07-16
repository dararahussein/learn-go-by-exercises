# Set 10 - Capstone: task tracker CLI and HTTP API

You are extending existing code, not starting from a blank file. Read all three
packages before changing anything and keep each change small enough to review.

Work in order: `01_store/task`, `02_http/api`, then `03_cli/cmd/task`. Each
earlier module becomes a real dependency of the later one.

Warm-up from memory: write the `errors.Is` branch and `httptest` setup you expect
the completion endpoint test to need, then verify those predictions in docs.

## Core work order

1. Make `go test ./set10/01_store/task` green: implement `Load`, `Add`, and `Pending`.
   Diagnose the planted `Complete` bug rather than replacing it blindly.
2. Run `go test -race ./set10/01_store/task`. HTTP handlers can call the store at the
   same time, so protect every read and write with the store's mutex.
3. Implement the CLI commands and use them against a temporary JSON file.
4. Make the provided HTTP GET test pass.
5. Write table-driven POST tests for valid JSON, malformed JSON, and empty text;
   only then implement POST. Add completion behavior with a 404 that checks
   `errors.Is(err, task.ErrNotFound)`.
6. Run `go fmt ./...`, `go vet ./...`, `go test ./...`, and
   `go test -race ./set10/...`.

The API should support:

- `GET /tasks` - pending tasks as JSON
- `POST /tasks` with `{"text":"learn Go"}` - create a task, status 201
- `POST /tasks/{id}/done` - complete a task, status 204 or 404

## Optional extensions

- `--file` with the `flag` package
- `list --all` and a remove command
- graceful shutdown with `signal.NotifyContext`
- a concurrent URL-check command, with a measured reason for concurrency

Expected time: 5-8 hours over the final one or two days.
