/*

From crashes to explicit failure — THE Go idiom.

In many languages, failure is an exception that silently climbs the stack.
In Go, a function that can fail says so in its signature by returning an extra
`error` value, and the CALLER must deal with it, visibly:

	value, err := mightFail()
	if err != nil {
		return 0, err   // propagate it upward
	}
	// ... use value

(If you did the Haskell course: `(T, error)` plays the role Maybe/Either did.
Failure lives in the type, not in a hidden channel.)

Tasks:
- [ ] Implement each function so it NEVER panics (replace the TODO stubs).
- [ ] Convention: when returning a non-nil error, return the ZERO value
      alongside it: `return 0, errors.New(...)`.
- [ ] Run `go test ./set02/` after every change.

Tools and techniques:
- errors.New("message") makes a simple error.
- fmt.Errorf("no age for %q", name) makes a formatted one (add the "fmt" import).
- Error messages are lowercase, no trailing punctuation (Go convention).

Expected time to complete: 60 minutes.

*/

package set02

import "errors"

// Original version that crashes:
//
//	func Head(xs []int) int { return xs[0] }  // panics on empty slice!
//
// Refactored to return an error — DONE FOR YOU:
func Head(xs []int) (int, error) {
	if len(xs) == 0 {
		return 0, errors.New("head of empty slice")
	}
	return xs[0], nil
}

// Original: func Tail(xs []int) []int { return xs[1:] }  // panics on empty!
func Tail(xs []int) ([]int, error) {
	return nil, nil // TODO: implement Tail (error on empty)
}

// Original: func Div(a, b int) int { return a / b }  // panics when b == 0!
func Div(a, b int) (int, error) {
	return 0, nil // TODO: implement Div (error on divide by zero)
}

// DivFirstTwo divides the first element of xs by the second.
// This one is about PROPAGATION: get the elements (too few? error), then call
// your Div and pass its error along if there is one. The
// `if err != nil { return 0, err }` dance is normal Go — embrace it.
func DivFirstTwo(xs []int) (int, error) {
	return 0, nil // TODO: implement DivFirstTwo
}

// FindEven returns the first even number in xs, or an error if there is none.
func FindEven(xs []int) (int, error) {
	return 0, nil // TODO: implement FindEven
}

// LookupAge returns ages[name], or an error whose message CONTAINS the missing
// name (use fmt.Errorf with %q). Future-you debugging at 2am will thank you for
// putting the data in the error message.
func LookupAge(ages map[string]int, name string) (int, error) {
	return 0, nil // TODO: implement LookupAge
}
