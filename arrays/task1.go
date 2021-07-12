package main

import "fmt"

func main() {
	var workArray [10]uint8
	var a, b, a1, b1 uint8
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Scan(&a, &b)
		a1 = workArray[a]
		b1 = workArray[b]
		//a1, b1 = b1, a1
		a1 = a1 + b1
		b1 = a1 - b1
		a1 = a1 - b1
		workArray[a] = a1
		workArray[b] = b1
	}
	for _, elem := range workArray {
		fmt.Print(elem, " ")
	}
}
