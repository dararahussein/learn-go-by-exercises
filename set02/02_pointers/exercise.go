/*

Pointers explain the difference between changing a copy and changing the
original. Go passes every argument by value. A pointer is a value that refers
to another value.

Tasks:
- [ ] Implement the TODOs.
- [ ] Run `go test`; the board advances from Increment to Counter.
- [ ] Before entering `04_debugging`, predict why a value-receiver method might
      change nothing.

Expected time: 45 minutes.

*/

package pointers

// DONE FOR YOU: write zero through p into the caller's variable.
func SetZero(p *int) { *p = 0 }

func Increment(p *int) {
	// TODO: add one to the value p points at
}

type Counter struct {
	n int
}

func (c Counter) Value() int {
	return 0 // TODO: return c.n
}

func (c *Counter) Inc() {
	// TODO: increment the original Counter
}
