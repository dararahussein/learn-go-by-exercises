# Set05 — Building a real project: packages, modules & tooling

Everything so far lived in one folder. Real projects are split into packages
across directories. This set is short on new syntax and long on the stuff you
need to actually *start a project*: how Go code is organized, imported, and run.

Expected time to complete: 45–60 minutes (mostly reading + one small exercise).

## The exercise

1. Implement `Greet` in [greeting/greeting.go](greeting/greeting.go).
2. Make the test pass:  `go test ./set05/...`
3. Run the program that imports your library:
   `go run ./set05/cmd/hello Ada`

That's it — the point is to *feel* one package importing another.

## The five things to understand

### 1. A module (go.mod)
The file `go.mod` at the repo root declares this whole repo as one **module**
named `learn-go-by-exercises`. A module is a unit you version and share. You
make one for your own project with:

```sh
go mod init github.com/you/yourproject
```

### 2. A package is a directory
Every `.go` file starts with `package X`. All files in one directory share the
same package. `package main` is special: it builds an executable with a
`func main()`. Any other name (`package greeting`) is a library.

### 3. Import paths = module path + directory
`greeting` lives at `set05/greeting/`, so its import path is
`learn-go-by-exercises/set05/greeting`. That's exactly what `cmd/hello/main.go`
imports. No magic — path = module name + folders.

### 4. Exported vs unexported = capitalization
`Greet` (capital G) is visible to other packages. `greet` (lowercase) would be
private to its own package. This is Go's ONLY access control. Name things
lowercase by default; capitalize only what you deliberately expose.

### 5. The conventional layout
You don't need this on day one, but you'll see it everywhere:

```
yourproject/
  go.mod
  main.go              # or cmd/<name>/main.go for multiple binaries
  cmd/                 # executables (one dir per binary)
  internal/            # private packages (Go FORBIDS importing these from
                       #   outside your module — great for app code)
  <somepkg>/           # exported library packages
```

Start flat (just `main.go`). Split into packages when a file gets unwieldy or a
chunk of code is clearly reusable. Don't over-engineer the structure early.

## The toolchain you'll use daily

| Command | What it does |
|---|---|
| `go run ./path` | compile + run a program |
| `go build ./...` | compile everything, report errors (no run) |
| `go test ./...` | run all tests |
| `go test -race ./...` | run tests with the data-race detector |
| `go vet ./...` | static checks for likely bugs |
| `go fmt ./...` | auto-format (the editor does this on save) |
| `go mod tidy` | add missing / drop unused dependencies in go.mod |
| `go doc fmt.Println` | read docs for any package or symbol in the terminal |
| `go get github.com/x/y` | add a third-party dependency |

`./...` means "this directory and everything under it."

## Starting your own project (do this right after Set06)

```sh
mkdir myapp && cd myapp
go mod init github.com/you/myapp
# create main.go with: package main / func main() { ... }
go run .
```

Then just grow it. Add a package when you need one. Add a dependency with
`go get`. You already know enough to build.
