package main

import "fmt"

func main() {
	var a, b uint
	var c float64
	fmt.Scan(&a, &b)
	c = (a + b) / 2
	fmt.Print(c)
}
