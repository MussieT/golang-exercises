package gomistakes

import (
	"fmt"
)

// Variable shadowing - a variable declared in a block can be redclared in an inner block.

type Person struct {
	name string
}

func VariableShadowing() {

	var dude Person

	if true {
		// The outer dude is shadowed
		dude := Person {
			name: "dude",
		}

		fmt.Println(dude)

	} 

	fmt.Println(dude.name)
}
