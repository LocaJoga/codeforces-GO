package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var word string
	fmt.Scan(&word)
	if IsUpper(word) == true {
		fmt.Println(strings.ToUpper(word))
	} else {
		fmt.Println(strings.ToLower(word))
	}
}

func IsUpper(s string) bool {
	lower := 0
	upper := 0
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			lower++
		} else {
			upper++
		}
	}
	if lower > upper || lower == upper {
		return false
	}
	return true
}
