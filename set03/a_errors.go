/*

Errors are ordinary values. A function that can fail says so in its signature,
and callers inspect or return that error explicitly.

Tasks:
- [ ] Implement the TODOs without panicking.
- [ ] Preserve causes with `%w`, not `%v`, when adding context.
- [ ] Run one test at a time with `go test -run TestName ./set03/`.

Expected time: 75 minutes.

*/

package set03

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

// DONE FOR YOU: a safe replacement for xs[0].
func Head(xs []int) (int, error) {
	if len(xs) == 0 {
		return 0, errors.New("head of empty slice")
	}
	return xs[0], nil
}

func Tail(xs []int) ([]int, error) {
	return nil, nil // TODO: error on an empty slice
}

func Div(a, b int) (int, error) {
	return 0, nil // TODO: error when b is zero
}

func DivFirstTwo(xs []int) (int, error) {
	return 0, nil // TODO: validate length, then call Div and propagate its error
}

// LookupAge should wrap ErrNotFound and mention name when the key is absent.
func LookupAge(ages map[string]int, name string) (int, error) {
	return 0, nil // TODO: use fmt.Errorf with %w
}

// TeamAge adds operation-level context while preserving LookupAge's cause.
func TeamAge(ages map[string]int, names []string) (int, error) {
	return 0, nil // TODO: sum ages; wrap any lookup error with the operation context
}

// Keep fmt available while the learner works through imports incrementally.
var _ = fmt.Errorf
