package main

import (
	"fmt"
	"unicode"
)

func latinCheck(sRune []rune) bool {
	for i := range sRune {
		if unicode.Is(unicode.Latin, sRune[i]) || unicode.IsDigit(sRune[i]){
			continue
		} else {
			return false
			break
		}
	}
	return true
}

func main() {
	var s string
	fmt.Scan(&s)
	sRune := []rune(s)
	if len(sRune) >= 5 && latinCheck(sRune) {
		fmt.Print("Ok")
	} else {
		fmt.Print("Wrong password")
	}
}