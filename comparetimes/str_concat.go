package comparetimes

import (
	"fmt"
	"os"
	"time"
)

func ConcatStr() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	elapsed := time.Since(start)
	fmt.Printf("str Concat elapse: %s\n", elapsed)
}