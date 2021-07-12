package main

import "fmt"

func main() {
	var a, b, c, sum int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b, &c)
		sum = b + c
		fmt.Print(sum, " ")
	}

}