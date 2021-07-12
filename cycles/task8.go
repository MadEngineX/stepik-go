package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		a, b, a2, b2 int
		a1, b1 string
	)

    fmt.Scan(&a, &b)
	a1 = strconv.Itoa(a)
	b1 = strconv.Itoa(b)
	a2 = len(a1)
	b2 = len(b1)

	for i := 0; i < a2; i++ {
		for j := 0; j < b2; j++ {
			if a1[i] == b1[j] {
				fmt.Print(string(a1[i]), " ")
			}
		}
	}

}
