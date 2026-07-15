/*

Pointers — the difference between changing a copy and changing the original.

Go passes arguments by VALUE (a copy). So a function that takes an `int` can't
change the caller's int. To let it change the original, you pass a POINTER — the
variable's address — and the function writes through it:

	&x   is "the address of x"        (type: *int)
	*p   is "the value p points at"   (dereference)

This is also why methods that MUTATE use a pointer receiver `(c *Counter)`:
the method needs the real struct, not a copy.

Tasks:
- [ ] Implement the stubs (marked TODO).
- [ ] Run `go test ./set02/` after every change.

Expected time to complete: 30 minutes.

*/

package set02

// DONE FOR YOU: write 0 into whatever p points at, changing the caller's value.
func SetZero(p *int) {
	*p = 0
}

// Increment adds 1 to the int that p points at.
func Increment(p *int) {
	// TODO: implement Increment
}

// Swap exchanges the values that x and y point at: after Swap(&a, &b),
// a holds b's old value and vice-versa.
func Swap(x, y *int) {
	// TODO: implement Swap
}

// Counter holds a count. n is lower-case, so it's unexported: only code in
// THIS package can touch it directly. Callers change it through methods.
type Counter struct {
	n int
}

// Value reads the count. A value receiver (c Counter) is fine for reading.
func (c Counter) Value() int {
	return 0 // TODO: return c.n
}

// Inc increases the count by 1. This MUST use a pointer receiver (c *Counter)
// — with a value receiver it would bump a copy and the caller would see no
// change. (It's already written as a pointer receiver for you.)
func (c *Counter) Inc() {
	// TODO: implement Inc
}
