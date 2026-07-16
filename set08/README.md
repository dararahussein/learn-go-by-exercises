# Set 08 - Races, synchronization, and cancellation

Start with the unsafe counter exactly as written:

Work in order: `01_races`, then `02_context`. Use `check` inside each
directory except for the explicit race demonstration below.

```sh
go test ./set08/01_races -args -race-demo
go test -race ./set08/01_races -args -race-demo
```

The first command may appear green. The second should report a data race. Save
the important parts of that report in `MISTAKES.md` and explain the conflicting
accesses. Leave `UnsafeCounter` as the isolated demonstration, then implement
`SafeCounter` and run the full set with `-race`. The demo is skipped unless the
explicit `-race-demo` flag is present, so a completed full suite can be clean.
A test passing once is not evidence that concurrent code is safe.

Warm-up from memory: draw the goroutines and channel ownership in `Merge` and
name the one goroutine allowed to close its output.

Expected time: 4-6 hours.

Optional stretch: compare `sync.Mutex`, `sync.RWMutex`, and a single owner
goroutine for the counter. Implement only one alternative and justify it.
