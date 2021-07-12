package main

import(
	"fmt"
	"strconv"
)

func main() {
	var a string
	fmt.Scan(&a)
	aRune := []rune(a)
	for i := range aRune {
		s, _ := strconv.Atoi(string(aRune[i]))
		fmt.Print(s*s)
	}
}