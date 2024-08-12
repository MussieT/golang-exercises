package comparetimes

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func StrJoinFnc() {
	start := time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))

	elapsed := time.Since(start)
	fmt.Printf("str Join elapse: %s\n", elapsed)
}
