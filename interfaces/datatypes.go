package interfaces

import "fmt"

func GetString(name interface{}) interface{} {
	var a interface{}
	var b interface{}

	a = 5
	b = a

	a = 10

	fmt.Printf("a is %d\n", a)
	fmt.Printf("b is %d\n", b)

	return name
}

