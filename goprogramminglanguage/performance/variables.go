package performance

import "fmt"

func TypeCoercision() {
	str := "hello"
	bytes := []byte(str)

	fmt.Println("%b", bytes)

	// To save additional memory allocation
	bytesTwo := []byte(str)[:len(str)]
	fmt.Println(bytesTwo)
}
