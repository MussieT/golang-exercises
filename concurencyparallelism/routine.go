package concurencyparallelism

import (
	"fmt"
	"time"
)

// In golang, each concurrently executing activities are called go routines
// New goroutines are created by the go statement.

// GOMAX-PROCS - is by default equal to the number of available CPU cores.

func GoRoutineSample() {
    go spinner(100 * time.Millisecond)

	const n = 45
	fibN :=  fib(n)
 	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}
