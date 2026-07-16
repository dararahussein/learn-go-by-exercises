/*

DEBUGGING — the other half of the job.

Every function below COMPILES and looks plausible, but its test FAILS. Unlike
the earlier modules, nothing here is a stub: the code is written, and wrong. Your
job is to read the failing test, form a theory about the bug, and fix the code.

How to work:
  1. `go test -v` - read what the current debugging step got vs. wanted.
  2. Predict the cause before you change anything. Write your guess in the
     `// A:` line under each function.
  3. Fix ONE thing, re-run. If it's still red, your theory was wrong — good,
     that's information.

These three bugs (off-by-one, early-return, int division) are the three you will
personally ship in your first month of Go. Meeting them here is cheap.

Expected time to complete: 30 minutes.

*/

package debugging

// BUG: LastIndexOf should return the index of the LAST occurrence of target,
// or -1 if absent. LastIndexOf([]int{10, 20, 10}, 10) should be 2.
// Q: Why does it return the wrong index? Look at the loop bound.
// A:
func LastIndexOf(xs []int, target int) int {
	last := -1
	for i := 0; i < len(xs)-1; i++ {
		if xs[i] == target {
			last = i
		}
	}
	return last
}

// BUG: Contains should report whether target appears anywhere in xs.
// Contains([]int{1, 2, 3}, 3) should be true.
// Q: What does `return false` inside the loop do too early?
// A:
func Contains(xs []int, target int) bool {
	for _, x := range xs {
		if x == target {
			return true
		}
		return false
	}
	return false
}

// BUG: Average should return the mean as a float64.
// Average([]int{1, 2}) should be 1.5, not 1.0.
// Q: `sum` and `n` are both int. When does the division happen — before or
//
//	after the conversion to float64?
//
// A:
func Average(xs []int) float64 {
	if len(xs) == 0 {
		return 0
	}
	sum, n := 0, len(xs)
	for _, x := range xs {
		sum += x
	}
	return float64(sum / n)
}

/*
RETRIEVAL — answer from memory before checking (edit the // A: lines):

Q1: Reading a key that isn't in a map (m["missing"]) does NOT panic. What does
    it return, and how do you tell "absent" from "present-but-zero"?
A1:

Q2: a := []int{1, 2, 3}; b := a[:2]; b = append(b, 99). What is a[2] now, and
    why? (Predict, then write a tiny program to check.)
A2:

Q3: What is the zero value of a []int, and is `append` to that zero value legal?
A3:
*/
