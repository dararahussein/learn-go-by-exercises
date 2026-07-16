package goroutines

import "sync"

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

// Parallel runs every function concurrently and preserves input order.
// Writing to distinct, preallocated slice indices does not require a mutex.
func Parallel(fns []func() int) []int {
	return nil // TODO
}

func SumConcurrent(groups [][]int) int {
	return 0 // TODO: one goroutine per group; collect partial sums through a channel
}
