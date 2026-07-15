/*

Goroutines — do many things at once.

Put `go` in front of a function call and it runs concurrently, in a lightweight
thread called a goroutine. The tricky part is WAITING for them and collecting
results safely. A sync.WaitGroup counts outstanding goroutines:

	var wg sync.WaitGroup
	for _, job := range jobs {
		wg.Add(1)                 // one more to wait for
		go func() {
			defer wg.Done()       // mark done when this goroutine returns
			work(job)
		}()
	}
	wg.Wait()                     // block until all Done

Rule of thumb: goroutines writing to DIFFERENT slice indices is safe; goroutines
writing to the SAME variable is a data race (that's the next two files).

Tasks:
- [ ] Implement the stubs (marked TODO).
- [ ] Run `go test ./set04/` after every change.

Expected time to complete: 45 minutes.

*/

package set04

import "sync"

// DONE FOR YOU: run every function concurrently and wait for all to finish.
func RunAll(fns []func()) {
	var wg sync.WaitGroup
	for _, f := range fns {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f()
		}()
	}
	wg.Wait()
}

// Parallel runs each fn concurrently and returns their results IN ORDER, so
// out[i] is the result of fns[i]. Trick: make the output slice up front with
// make([]int, len(fns)), then each goroutine writes to its OWN index out[i]
// (safe, no lock needed). Use a WaitGroup to wait for them all.
func Parallel(fns []func() int) []int {
	return nil // TODO: implement Parallel
}

// SumConcurrent sums each sub-slice in its own goroutine, then returns the grand
// total. A channel is a clean way to collect the partial sums:
//
//	partials := make(chan int)
//	for _, xs := range xss { go func() { partials <- sumOf(xs) }() }
//	// then receive len(xss) values and add them up
func SumConcurrent(xss [][]int) int {
	return 0 // TODO: implement SumConcurrent
}
