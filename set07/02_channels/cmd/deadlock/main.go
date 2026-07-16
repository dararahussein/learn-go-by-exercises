// Run this program separately: go run ./set07/02_channels/cmd/deadlock
package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 42 // BUG: who can receive while main is blocked here?
	fmt.Println(<-ch)
}

// Q: Copy the important line from the runtime dump and explain the wait cycle.
// A:
