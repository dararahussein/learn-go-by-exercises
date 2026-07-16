/* DEBUGGING: this function compiles and reports success, but loses the result. */
package debugging

import "strconv"

// Q: Which port and err exist inside the if block? Which ones are returned?
// A:
func ParsePort(text string) (int, error) {
	var port int
	var err error
	if text != "" {
		port, err := strconv.Atoi(text) // BUG: shadows the outer variables
		_ = port
		_ = err
	}
	return port, err
}

/*
RETRIEVAL:
Q1: Why is `fmt.Errorf("lookup: %w", err)` different from using `%v`?
A1:
Q2: If a method has a pointer receiver, do T, *T, or both satisfy an interface
    requiring that method? Verify with a compile-time interface check.
A2:
*/
