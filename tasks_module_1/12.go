package main

import (
	 "fmt"
	 "math"
)

func main() {
	var a, n, i float64
	fmt.Scan (&n)
	a = 2

	for i = 0; math.Pow(a, i) <= n; i++ {
		if math.Pow(a, i) <= n {
			fmt.Print(math.Pow(a, i), " ")
		}
	}
}
