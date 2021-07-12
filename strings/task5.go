package main

import(
	"fmt"
	"strings"
)
func main() {
	var s string
	fmt.Scan(&s)
	sRune := []rune(s)
	for i := range sRune {
		if strings.Count(s, string(sRune[i])) > 1 {
			s = strings.Replace(s, string(sRune[i]), "", -1)
		}
	}
	fmt.Print(s)
}