package datatypes

import (
	"fmt"
	"unicode/utf8"
)

func UtilizeStrings() {
	s := "hello world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	// Unicode runes..
	str := "cafê"

	fmt.Printf("%s, length %d\n", str, len(str))
	
	// for the characters length
	charCount := utf8.RuneCountInString(str)
	fmt.Printf("count: %d\n", charCount)

	// subsctrings..
	substr := str[:2]
	fmt.Printf("%s \n", substr)
}

// Immutability means that it is safe for two copies of a string to share the same underlying memory;
// making it cheap to copy strings of any length. Similarly, a string s and a substring like s[:7] may safely share
// the same data, so the substring operation is cheap.
