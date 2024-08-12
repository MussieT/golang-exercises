package datatypes

import (
	"fmt"
	"math"
)

// the type rune is similar to int32
// if the number is always positive .. just use uint. (but not related with performance - they are just 32 bit)
// var total_cities_in_ethiopia uint

func UseIntegers() {
	// total_cities_in_ethiopia = 9
	// fmt.Printf("%d", total_cities_in_ethiopia)

	// 
	fmt.Printf("min - %v, max - %v\n", math.MinInt32, math.MaxInt32)
}

// we can check integer overflow
func IntegerOverflow(counter int) int {
	if counter == math.MaxInt {
		panic("int overflow")
	}
	return counter + 1
}

func FloatingPoints() {
	fmt.Printf("smallest, %v largetst: %v\n", math.SmallestNonzeroFloat64, math.MaxFloat64)

	// NaN Equality
	// fmt.Printf("Nan equality %v\n", math.NaN() == math.NaN()) no value is equal to NaN, not even NaN itself (SA4012)go-staticcheck

	// wHhen comparing floating point numbers, compare there difference is in acceptable range. - because golang can only handle what it stores in its 62bit representation.
}

