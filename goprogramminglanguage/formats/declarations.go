package formats

import "fmt"

// Use camelCase (General Rule)

var name string = "mussie"
var i, j, k int
var lastName = new(string) // this is pointer
// can not be used in global scope - short declaration
// shortDeclaration := "short declaration"

func Declaration() {
	shortDeclaration := "short declaration"
	fmt.Printf("%s", shortDeclaration)
	fmt.Println()
	fmt.Printf("last %s", *lastName)
	fmt.Println()

	x := 10
	y := 8

	x, y = y, x

	fmt.Printf("x is %d and y is %d \n", x, y)
}
