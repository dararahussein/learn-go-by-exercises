/*

Pointers explain the difference between changing a copy and changing the
original. Go passes every argument by value. A pointer is a value that refers
to another value.

Tasks:
- [ ] Implement the TODOs.
- [ ] Run `go test -run 'TestIncrement|TestCounter' ./set02/`.
- [ ] Before fixing z_fixme.go, predict why its method changes nothing.

Expected time: 45 minutes.

*/

package set02

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

// TimestampedTemperature demonstrates composition through embedding. The
// fields and methods of Temperature are promoted to this type.
type TimestampedTemperature struct {
	Temperature
	RecordedAt string
}
