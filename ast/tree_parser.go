package main

import (
	"strings"
	"unicode"
)

func main() {
	input := "key [ subkey ] [ value ] [ ]"
	input2 := removeWS(input)

	for i := 0; i <= len(input2)-1; i++ {
		if input2[i] != 14 {

		}
	}

	println(input2)
	//for i, j := range input {
	//
	//}
}

func removeWS(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}
