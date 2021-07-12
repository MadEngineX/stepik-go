package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Scan(&a)
	aRune := []rune(a)
	for i := range aRune {
		if i < len(aRune) - 1  {
			b = b + string(aRune[i]) + "*"
		} else {
			b = b + string(aRune[i])
		}
	}
	fmt.Print(b)
}
