/*

This is a LIBRARY package (not `package main`), living in its own directory.
Its directory name and package name are both `greeting`. Other packages import
it by its full path:  learn-go-by-exercises/set05/greeting

Only identifiers that start with a Capital letter are EXPORTED (visible to other
packages). `Greet` is exported; a lowercase helper would be private to this
package. That capitalization rule IS Go's access control — there's no
public/private keyword.

Task:
- [ ] Implement Greet, then run the program that uses it:
      go run ./set05/cmd/hello Ada
- [ ] Run the test:  go test ./set05/...

*/

package greeting

// Greet returns "Hello, <name>! Welcome to Go."
// If name is empty, use "friend": "Hello, friend! Welcome to Go."
func Greet(name string) string {
	return "" // TODO: implement Greet
}
