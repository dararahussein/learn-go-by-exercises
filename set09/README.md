# Set 09 - Work in an existing codebase

Do not start by editing. Treat this directory like a repository you just joined.

Work in order:

1. `01_packages/greeting` - package/import warm-up
2. `02_codebase` - navigate and modify a multi-package application
3. `03_generics/collections` - focused generics exposure

Warm-up from memory: list the commands for one targeted test, all tests, static
checks, formatting, and the race detector without looking at the root README.

## 1. Trace before changing

Run `go run ./set09/02_codebase/cmd/bookshop 978-0-13-419044-0`, then write the call path
from `main` to the final price in `NOTES.md`. Identify which package owns book
data, normalization, and price policy. Use `go doc` on at least two symbols.

## 2. Diagnose a cross-package bug

`go test ./set09/02_codebase/catalog` stops at `01_FindReturnsStoredBook`. The symptom is
in the test; the cause is inside a loop in another file. Add temporary logging
if useful, explain the copy in `MISTAKES.md`, and make the smallest fix.

## 3. Add a feature test-first

Complete `Search` in `catalog/search.go`. Replace the deliberate failure in its
test with cases for title, author, case-insensitivity, no match, and empty query
before implementing it. Keep the API small.

## 4. Read and extend generic code

Use the provided `Filter[T]`, then implement `Map[T, U]`. This is exposure, not
a generics deep dive.

## 5. Practice module maintenance

Replace one manual struct comparison in an editable test with `cmp.Diff`:

```sh
go get github.com/google/go-cmp/cmp
go mod tidy
go test ./set09/...
```

Inspect the `go.mod` and `go.sum` changes. Keep them only if the dependency makes
the test clearer; explain your judgment. This is the one intentional third-party
dependency exercise in the course.

Expected time: 4-6 hours.

Optional stretch: add an HTTP read endpoint over Catalog without moving data
ownership into the handler package.
