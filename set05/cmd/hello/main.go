/*

This is the PROGRAM (package main) that ties into the greeting library.
It lives under cmd/hello/ — the conventional home for an executable named
"hello". A repo can have several: cmd/server, cmd/worker, etc.

Notice the import path: "learn-go-by-exercises/set05/greeting". That's the
module path (from go.mod) + the package's directory. This is the whole story
of how Go finds your other packages.

Run it:
    go run ./set05/cmd/hello           -> Hello, world! Welcome to Go.
    go run ./set05/cmd/hello Ada       -> Hello, Ada! Welcome to Go.

(It won't print correctly until you implement Greet in the greeting package.)

*/

package main

import (
	"fmt"
	"os"

	"learn-go-by-exercises/set05/greeting"
)

func main() {
	name := "world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(greeting.Greet(name))
}
