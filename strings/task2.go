package main

import (
	"fmt"
)

func palindrom(text string) bool {
	textRunes := []rune(text)
	for i := 0; i < len(textRunes)/2; i++ {
		if textRunes[i] != textRunes[len(textRunes) - 1 - i] {
			return false
		}
	}
	return true
}

func main() {
	var text string
	fmt.Scan(&text)
	if palindrom(text) {
		fmt.Print("Палиндром")
	} else {
		fmt.Print("Нет")
	}
}
