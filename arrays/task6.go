package main

import "fmt"

func main() {
	var array [100]int
	var a, n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	count := 0
	for i := 0; i<n; i++ {
		if array[i] > 0 {
			count++
		}
	}
	fmt.Print(count)

}
