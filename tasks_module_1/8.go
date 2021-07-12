package main

import "fmt"

func main() {
	var n, a, count, min int
	var arr []int
    count = 1
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}

	min = arr[0]
	for i := 0; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	for i := 0; i < n; i++ {
		if arr[i] == min {
			count++
		}
	}

	fmt.Print(count)
}
