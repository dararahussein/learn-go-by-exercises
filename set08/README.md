# Set 08 - Races, synchronization, and cancellation

Start with the unsafe counter exactly as written:

```sh
go test -run TestUnsafeCounterRace ./set08/
go test -race -run TestUnsafeCounterRace ./set08/
```

The first command may appear green. The second should report a data race. Save
the important parts of that report in `MISTAKES.md`, explain the conflicting
accesses, then protect the counter. A test passing once is not evidence that
concurrent code is safe.

Warm-up from memory: draw the goroutines and channel ownership in `Merge` and
name the one goroutine allowed to close its output.

Expected time: 4-6 hours.

Optional stretch: compare `sync.Mutex`, `sync.RWMutex`, and a single owner
goroutine for the counter. Implement only one alternative and justify it.
