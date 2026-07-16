/*

Tasks:
- [ ] Implement each function by replacing its zero-value stub (marked TODO).
- [ ] Run `go test` after every change. The board reveals one step at a time.

How this course works:
- Every function starts returning a dummy zero value so the file compiles and
  its test FAILS. Your job is to make the test pass.
- Go has no ternary and no `while` — `if`, `for`, and `switch` do everything.
- A function can return multiple values: `func f() (int, int)` ... `return a, b`.
- A Tour of Go, "Flow control": https://go.dev/tour/flowcontrol/1

Expected time to complete: 45 minutes.

*/

package functions

// DONE FOR YOU — this is the shape every exercise takes.
// func <name>(<params>) <return type> { ... }
func Double(x int) int {
	return x * 2
}

// Abs returns x without its sign: Abs(-3) == 3, Abs(3) == 3.
func Abs(x int) int {
	return 0 // TODO: implement Abs
}

// Max returns the larger of x and y.
// (Yes, Go 1.21+ has a built-in `max` — write it yourself anyway, with `if`.)
func Max(x, y int) int {
	return 0 // TODO: implement Max
}

// Clamp limits x to the range [lo, hi]:
// Clamp(5, 0, 10) == 5, Clamp(-5, 0, 10) == 0, Clamp(15, 0, 10) == 10.
func Clamp(x, lo, hi int) int {
	return 0 // TODO: implement Clamp
}

// DivMod returns both the quotient and the remainder of a / b.
// Your first function with two return values.
func DivMod(a, b int) (int, int) {
	return 0, 0 // TODO: implement DivMod
}

// SumTo returns 1 + 2 + ... + n. Use a `for` loop:
// for i := 1; i <= n; i++ { ... }
func SumTo(n int) int {
	return 0 // TODO: implement SumTo
}

// FizzBuzz returns "Fizz" for multiples of 3, "Buzz" for multiples of 5,
// "FizzBuzz" for multiples of both, and the number itself (as a string) otherwise.
// strconv.Itoa(n) turns an int into a string — you'll need to add an
// import ("strconv") block at the top (your editor can do it on save).
func FizzBuzz(n int) string {
	return "" // TODO: implement FizzBuzz
}
