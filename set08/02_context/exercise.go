package contextx

import (
	"context"
	"time"
)

// CountUntilCanceled sends increasing integers at interval and closes its
// output after ctx is canceled. It must not leave a goroutine blocked on send.
func CountUntilCanceled(ctx context.Context, interval time.Duration) <-chan int {
	out := make(chan int)
	close(out) // TODO: replace with a goroutine selecting on ticker.C and ctx.Done()
	return out
}
