package datatypes

import (
	"encoding/json"
	"fmt"
	"runtime"
)

func SlicesPractice() {
	// months := [...]string{1: "January", 2: "February"}
	// fmt.Println(months)

	// Slice is backed by an underlying array data structure

	// There is a slice length, and capacity
	s1 := make([]int, 3, 6) // in this case the slice length is 3 and capacity is 6 

	fmt.Println("first slice is: ", s1) // this will have - 0 values for the first 3 indeces
	fmt.Println("slice length? ", len(s1)) // it will print the available elements length.
	fmt.Println("slice capacity? ", cap(s1)) // it will print the available elements length.
	// fmt.Printf("5th element: %d", s1[4]) - it will return error if we try to access 5th index for example, even if it has the capacity, we haven't used it yet.

	s2 := s1[1:3]

	fmt.Println("s2 is: ", s2) // s2 will be pointing to the same underlying array starting from index 1.
	fmt.Println("slice length two? ", len(s2))

	s2 = append(s2, 4, 4, 4, 3, 2) // this isn't gonna return an error but create a new underlying array double the capacity (fist s2 capacity - 5) and s2 will be pointing to it

	fmt.Println("modified s2: ", s2)
	fmt.Println("slice s2 after modification: ", len(s2))

	// TLDR: Slice length is the number of available elements in elements in the slice, whereas slice capacity is the number of elements in the backing array
}

// Benchmark 0 capacity slice, append function and accessing index in slices.
func SliceWithZeroLength() {

	s3 := make([]int, 0) // an array of 0 length
	fmt.Println("s3 is: ", s3)

}

func log(i int, s []string) {
	fmt.Printf("%d: empty=%t\t nil=%t\n", i, len(s) == 0, s == nil)
}

// We should favor returning nil (Option 1 and 2) slice instead of empty (allocated) slice - Option - 3, because well the second one allocates space in memorys. 
func NilVsEmptySlice() {
	var s []string // Option 1
	log(1, s)

	s = []string(nil) //  Option 2 syntactic sugar - more readable
	log(2, s)

	// s_2 := []string()
	// log(3, s_2)

	s = make([]string, 0) // Option 3 this one allocates memory
	log(3, s)

	// Marshaling nil and empty
	type Customer struct {
		ID string
		Operations []float32
	}
	var s1 []float32

	customer_1 := Customer {
		ID: "foo",
		Operations: s1,
	}

	b, _ := json.Marshal(customer_1)
	fmt.Println(string(b))

	s2 := make([]float32, 0)
	customer_2 := Customer {
		ID: "foo",
		Operations: s2,
	}

	c, _ := json.Marshal(customer_2)
	fmt.Println(string(c))
}

// We can check both nil and empty slices by their length since it is 0
func CheckingNilEmptySlices() {
	s := make([]int, 0)
	fmt.Printf("empty_slice_length - %d\n", len(s))
	s2 := []string(nil)
	fmt.Printf("\nnil_slice_length - %d\n", len(s2))
}

func CopySlices() {
	src := []int{1, 2, 3}
	// dst := make([]int, len(src))
	// copy(dst, src)
	dst := append([]int(nil), src...)
	fmt.Printf("copied: %v\n", dst)
}

func UnexpectedSideEffects() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:2]
	// s3 := append(s2, 10)

	fmt.Printf("s1 is %v\n", s1)
	fmt.Printf("s2 is %v\n", s2)
}

func LargeSlice() {
	// This will take memory a lot.
	msg := make([]byte, 1000000)

	// this will slice it but it would have 1000000 capacity, it leads to memory leaks. (The new go will garbage collect this.)
	// msgSliced := msg[:5]

	// memory efficient - use 5 - capacity
	msgSliced2 := make([]byte, 5)
	copy(msgSliced2, msg)

	// msgType := getMessageType(msg)
	fmt.Println(cap(msgSliced2))
}

func getMessageType(msg []byte) []byte {
	return msg[:5]
}

// Garbage collect.
func CheckMemory() {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc / 1024)
}
