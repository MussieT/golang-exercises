package defersample

import "fmt"

func DeferSample() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
