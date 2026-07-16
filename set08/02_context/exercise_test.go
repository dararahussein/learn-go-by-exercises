package contextx

import (
	"context"
	"testing"
	"time"
)

func TestCountUntilCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	values := CountUntilCanceled(ctx, time.Millisecond)
	select {
	case _, ok := <-values:
		if !ok {
			t.Fatal("channel closed before producing a value")
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("worker did not produce a value")
	}
	cancel()
	deadline := time.After(100 * time.Millisecond)
	for {
		select {
		case _, ok := <-values:
			if !ok {
				return // an already-selected tick may be in flight; closure is the contract
			}
		case <-deadline:
			t.Fatal("worker leaked after cancellation")
		}
	}
}

/*
Q1: Why should context.Context normally be the first parameter and not stored
    in a long-lived struct?
A1:
Q2: Which operation in CountUntilCanceled could block forever unless the send
    itself participates in a select with ctx.Done()?
A2:
*/
