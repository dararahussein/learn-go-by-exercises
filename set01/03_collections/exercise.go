/*

Slices and maps — the two collections you'll use in every Go program.

Tasks:
- [ ] Implement each function by replacing its zero-value stub (marked TODO).
- [ ] Run `go test` after every change.

The four patterns to internalize (they cover ~90% of real Go data wrangling):
  1. build a slice:   var out []T; for ... { out = append(out, x) }; return out
  2. build a map:     m := make(map[K]V); for ... { m[k]++ }  (missing key == zero value)
  3. comma-ok:        v, ok := m[key]   // ok reports whether key was present
  4. deterministic:   maps range in RANDOM order — collect keys, sort, then range

Expected time to complete: 45 minutes.

*/

package collections

// DONE FOR YOU — pattern #1. Start empty, append in a loop, return.
// This exact shape is the most common thing you'll write in Go.
func Doubled(xs []int) []int {
	var out []int
	for _, x := range xs {
		out = append(out, x*2)
	}
	return out
}

// ArrayTotal is brief array exposure. [3]int has its length in its type;
// arrays are useful occasionally, but slices are the everyday collection.
func ArrayTotal(xs [3]int) int {
	return 0 // TODO: range over the array
}

// Sum returns the total of xs. Sum(nil) == 0.
func Sum(xs []int) int {
	return 0 // TODO: implement Sum
}

// Evens keeps only the even numbers of xs, in order. (Filtering = pattern #1
// with an `if` around the append.)
func Evens(xs []int) []int {
	return nil // TODO: implement Evens
}

// IndexOf returns the index of the first target in xs, or -1 if absent.
func IndexOf(xs []int, target int) int {
	return 0 // TODO: implement IndexOf
}

// WordCount counts occurrences of each word — pattern #2.
// WordCount([]string{"go","is","go"}) == map[string]int{"go":2, "is":1}
func WordCount(words []string) map[string]int {
	return nil // TODO: implement WordCount
}

// Lookup wraps the comma-ok idiom (pattern #3): return the value and whether
// the key was present. This is Go's built-in "Maybe".
func Lookup(m map[string]int, key string) (int, bool) {
	return 0, false // TODO: implement Lookup
}

// SortedKeys returns the map's keys in alphabetical order — pattern #4, the fix
// for map randomness. Collect keys into a []string, then sort.Strings(keys)
// (add an import ("sort") block for that).
func SortedKeys(m map[string]int) []string {
	return nil // TODO: implement SortedKeys
}
