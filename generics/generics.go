package generics

import "fmt"

// What is generic - its a golang feature which allows you to write reusable and type-safe code.
// TODO: generics - use cases in real problems ..

// Before generics - two functions
func SumInts(m map[string]int64) int64 {
	var s int64

	for _, v := range m {
		s += v
	}

	return s
}


func SumFloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}

	return s
}

func ShowResultsBeforeGeneric() {
	ints := map[string]int64 {
		"first": 21,
		"second": 54,
	}

	floats := map[string]float64 {
		"first": 78.3,
		"second": 55.5,
	}

	fmt.Printf("Non generic - ints sum : %v floats sum: %v\n", SumInts(ints), SumFloats(floats))
}

// Generic
func SumIntsOrFloats[K comparable, V int64 | float64] (m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}

	return s
}

// Restructure
// TODO: generics - use cases in real problems ..

type Number interface { int64 | float64 }

func RestructuredSum[K comparable, V Number] (m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}

	return s
}

func RestructuredPrintSums() {
	ints := map[string]int64 {
		"first": 21,
		"second": 54,
	}

	floats := map[string]float64 {
		"first": 78.3,
		"second": 55.5,
	}

	fmt.Printf("generic -  %v,  type inferred %v, ", SumIntsOrFloats(ints), SumIntsOrFloats[string, float64](floats))
	fmt.Printf("with interface - %v\n", RestructuredSum(floats))
}
