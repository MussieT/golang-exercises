package datatypes

import "fmt"

// All memebers must be homogeneus. Fixed size.
func ArrayManipulation() {
	// Assign 99th value to -1
	arr := [...]int{99: -1}
	
	fmt.Printf("%d", arr)
}
