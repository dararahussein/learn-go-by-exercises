# Farewell — you can build things in Go now

If you finished Set06, you've written, end to end, the skeleton most small Go
programs share: typed data (`struct`), behavior (`method`/`interface`),
explicit errors (`(T, error)`), collections, JSON, files, an HTTP handler,
concurrency, and a multi-package layout. That's the "get by on a small project"
bar. Time to build.

## Do this now: start your project

Seriously, today. The fastest way to cement all of this is to build something
you want. Pick something small:

- a CLI that does one useful thing (rename files, scrape a number, track habits)
- a JSON HTTP API with 2–3 endpoints (extend the Set03 server + Set06 store)
- a concurrent fetcher (a list of URLs → status codes, using Set04)

```sh
mkdir myapp && cd myapp
go mod init github.com/you/myapp
go run .
```

## When you get stuck (you will)

- **[pkg.go.dev/std](https://pkg.go.dev/std)** — the standard library. This is
  where most answers live. Learn to skim a package page fast.
- **[go.dev/tour](https://go.dev/tour/)** — the official tour; good for filling
  gaps in syntax you skated over.
- **[gobyexample.com](https://gobyexample.com)** — a runnable example for
  nearly every feature. Great for "how do I X again?"
- **`go doc <thing>`** in the terminal — instant docs, offline.

## Worth learning next (in rough priority)

1. **`context.Context`** — cancellation and deadlines; every real server and
   client uses it. (Natural follow-up to Set04.)
2. **Error wrapping** — `fmt.Errorf("...: %w", err)`, `errors.Is`, `errors.As`.
   You met the basics; this is the production version.
3. **Generics** — `func Map[T, U any](...)`; the `slices` and `maps` packages.
4. **A web framework or router** — the stdlib `net/http` is enough for a long
   time; reach for `chi` or `echo` only when you feel the need.
5. **Database access** — `database/sql` + a driver, or `sqlc` / `sqlx`.
6. **[Effective Go](https://go.dev/doc/effective_go)** and
   **[Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)** — how
   Go is meant to be written. Short and worth it once the syntax is automatic.

## The one habit that matters

Keep the feedback loop tight: editor + `go test`/`go run` side by side, small
changes, run constantly. It's how you learned here, and it's how you'll keep
getting better. Go build something. 🚀
