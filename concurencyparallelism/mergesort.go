package concurencyparallelism

import "sync"

func parallelMergesortV1(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		parallelMergesortV1(s[:middle])
	}()

	go func() {
		defer wg.Done()
		parallelMergesortV1(s[middle:])
	}()

	wg.Wait()
	merge(s, middle)

}

func merge(s []int, m int) {
	// whaterver 
}

// If the workload that we want to parallelize is too small, meaning we’re going to
// compute it too fast, the benefit of distributing a job across cores is destroyed: the time
// it takes to create a goroutine and have the scheduler execute it is much too high com-
// pared to directly merging a tiny number of items in the current goroutine. Although
// goroutines are lightweight and faster to start than threads, we can still face cases
// where a workload is too small.
