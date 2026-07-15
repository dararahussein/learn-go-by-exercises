# Set 07 - Goroutines, channels, and select

Run targeted tests while building. `Merge` is intentionally the hardest task;
complete `Relay` first because it contains half of the same shape.

Warm-up from memory: explain why an HTTP client must close a response body even
after it has read all the bytes.

For the deadlock lab, run `go run ./set07/cmd/deadlock`, read the runtime dump,
write the cause in `cmd/deadlock/main.go`, then repair it. Do not reach for
concurrency automatically: the reflection asks when a plain loop is better.

Expected time: 4-6 hours.

Optional stretch: build a three-stage generate -> square -> sum pipeline and
state which goroutine owns closing each channel.
