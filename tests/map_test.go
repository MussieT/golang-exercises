package datatypes

import (
	"fmt"
	"testing"
)

func MapWithInterface() {
	m := map[int]interface{}{}
	for i := 1; i <= 100; i++ {
		m[i] = nil
	}
}

func MapWithEmptyStruct() {
    m := map[int]struct{}{}
    for i := 1; i <= 100; i++ {
        m[i] = struct{}{}
    }
}


func Benchmark_Interface(b *testing.B) {
	fmt.Printf("testing n %d\n", b.N)
    for i := 0; i < b.N; i++ {
        MapWithInterface()
    }
}

func Benchmark_EmptyStruct(b *testing.B) {
    for i := 0; i < b.N; i++ {
        MapWithEmptyStruct()
    }
}