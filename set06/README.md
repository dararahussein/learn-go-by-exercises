# Set04 — Capstone: a task-tracker CLI

This is the payoff. You build a small but *complete* program: a command-line
todo list that saves to a JSON file. Every piece uses something you've already
practiced, so nothing here is new — it's assembly.

## What you're building

```
$ go run ./set06 add "learn go"
added #1: learn go
$ go run ./set06 add "build a project"
added #2: build a project
$ go run ./set06 list
#1 learn go
#2 build a project
$ go run ./set06 done 1
completed #1
$ go run ./set06 list
#2 build a project
```

State lives in `tasks.json` next to where you run it.

## The two files

- **store.go** — the logic: a `Task` struct, a `Store`, and methods
  (`Add`, `Complete`, `Pending`, `Save`, `Load`). This is unit-tested.
- **main.go** — the CLI: reads `os.Args`, calls the store, prints results.

## Work order

1. **Green the tests first.** Implement the methods in `store.go` until
   `go test ./set06/` passes. The tests pin down the exact behavior.
2. **Then build the CLI.** Implement the `add` / `list` / `done` commands in
   `main.go`. There's no test for `main` — you verify it by *using* it with
   `go run ./set06 ...`, which is exactly how you'll verify your own projects.

## When it works

You've now written, end to end, the same skeleton that most small Go tools
share: typed data, methods, error handling, JSON persistence, a CLI. That's
the "get by on a small project" bar from the README — you're there.

Ideas to make it yours (great first self-directed changes):
- a `rm <id>` command (slice deletion)
- show completed tasks with a `✓`, in a `list --all`
- a `--file` flag to pick the JSON path (see the `flag` package)

Then read `C_FAREWELL.md` and go start something real.
