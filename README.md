# Learn Go by Exercises

A hands-on course of small, focused exercises with instant test feedback.
Modeled on `learn-haskell-by-exercises`: you learn by *writing code* and getting
fast PASS/FAIL feedback, not by reading long tutorials.

## Goal

Get you writing real Go fast — enough to start building a small CLI or web tool
of your own. Six short sets; roughly a set a day, comfortably under a week. The
last set IS a small project, and by then you'll have built every piece of it
once already.

## How it works

Every exercise file has functions that start as stubs returning a dummy value,
so the file **compiles but its test fails**. Your job is to turn red tests
green.

```sh
go test ./set01/        # see the failing "board" for a set
go test ./set01/ -v     # verbose: every test named PASS/FAIL
go test -run TestAbs ./set01/   # just one
```

The first function in each file is **done for you** — study it, it shows the
pattern the rest of the file wants. A `// TODO` marks what's yours to write.

## The feedback loop (the most important habit)

1. Open an exercise file; read the header comment (your task list).
2. Replace a `// TODO` stub with a real implementation. Save.
3. `go test ./setNN/` — read the failure, adjust, repeat until green.

Keep the terminal beside your editor. Install the **VS Code Go extension** (or
your editor's `gopls` integration) so you see errors *as you type* — that's your
fast loop, the equivalent of a file-watcher. If you have to alt-tab to see
errors, your loop is too slow and you'll learn slower.

## Getting started

```sh
go version          # need Go 1.24+ (any recent Go is fine)
go test ./set01/    # should print FAILURES — that's the starting line
```

Then start turning `set01` green.

## The six sets

Work through them in order; within a set, go alphabetically (`a_`, `b_`, ...).

| Set | Theme | You learn |
|-----|-------|-----------|
| **set01** | Fundamentals | functions, control flow, strings & runes, slices, maps |
| **set02** | Types & behavior | structs, methods, errors `(T, error)`, interfaces, pointers |
| **set03** | Standard library | JSON, files, a small HTTP server |
| **set04** | Concurrency | goroutines, channels, `select`, `sync.Mutex`, the race detector |
| **set05** | Real projects | packages, modules, imports, the `go` toolchain |
| **set06** | Capstone | build a task-tracker CLI that uses all of the above |

Some runnable programs are launched directly, e.g.:

```sh
go run ./set03/c_http          # the HTTP server
go run ./set05/cmd/hello Ada   # cross-package program
go run ./set06 add "learn go"  # the capstone CLI
```

## Rules of engagement

- Replace `// TODO` stubs with your implementation. Don't edit the `_test.go`
  files (they define what "correct" means) or the done-for-you examples.
- Questions embedded as `// Q:` comments want written answers — edit the file
  and fill in the `// A:` line. Writing it down is how you find out whether you
  actually understand it.

## When you're stuck

- Give it an honest go, jot your thoughts as a comment, and move on — come back
  after the next exercise.
- Ask an LLM to *explain the concept* (not solve the exercise — the repo's
  `AGENTS.md` tells it not to). Pasting a compiler error verbatim gets great
  explanations.
- Two references you'll use constantly:
  - <https://pkg.go.dev/std> — the standard library (Go's Hoogle)
  - <https://gobyexample.com> — short runnable examples of every feature

## After you finish

See [set06/C_FAREWELL.md](set06/C_FAREWELL.md) — a curated "what next" list. The
short version: start your own project immediately. You'll learn more in the
first week of building than in any tutorial.
