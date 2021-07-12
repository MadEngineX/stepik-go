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
	for i := 0; i<n; i+=2 {
		fmt.Print(array[i], " ")
	}

}
