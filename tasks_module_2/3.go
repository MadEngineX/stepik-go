package main

import(
	"fmt"
)

func main() {
	var a, max string
	max = "0"
	fmt.Scan(&a)
	aRune := []rune(a)
	for i := range aRune {
		if string(aRune[i]) > max {
			max = string(aRune[i])
		}
	}
	fmt.Print(max)
}
