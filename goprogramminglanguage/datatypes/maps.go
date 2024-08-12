package datatypes

import "fmt"

func MapFunctions() {
	m1 := map[string]int {
		"mussie": 44,
	}

	
	m1["sami"] = 55
	fmt.Printf("%d\n", m1["sami"])

}
