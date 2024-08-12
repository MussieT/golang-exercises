package formats

import (
	"flag"
	"fmt"
)

func PointerThings() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

}

func incr(p *int) int {
	*p++
	return *p
}

// Flag
func FlagImplementation() {
	name := flag.String("name", "John Doe", "Your name")
	count := flag.Int("count", 1, "Number of times to repeat")

	// Parse flags
	flag.Parse()

	// Access flag values
	fmt.Println("Name:", *name)
	fmt.Println("Count:", *count)
}
