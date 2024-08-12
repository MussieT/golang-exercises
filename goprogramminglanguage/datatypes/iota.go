package datatypes

import "fmt"

const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// shifting bits
const (
	_ = iota
	KB = 1 << (10 * iota) // here iota is 1. so, 10 * 1 = 10, with the left shift operator it is: 2 the power of 10 which is 1024!
	MB
	GB
	TB
	PB
)

func PrintIotas() {
	fmt.Printf("%d\n", MB)

	var s uint = 5
	i := 1 << s
	fmt.Printf("%d\n", i)
}
