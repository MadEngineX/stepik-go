package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, a1 int
	var res string
	fmt.Scan(&a, &a1)
	b := strconv.Itoa(a)
	b1 := strconv.Itoa(a1)
	for i := 0; i < len(b); i++ {
		if string(b[i]) != b1 {
			res = res + string(b[i])
		}
	}
	fmt.Print(res)
}
