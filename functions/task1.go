package main

import "fmt"

func main() {
	fmt.Print(minimumFromFour())
}
func minimumFromFour() int {
	var a, min, n int
	var arr []int
	n = 4

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
    return min
}






