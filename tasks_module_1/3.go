package main

import "fmt"

func main() {
	var k uint32
	fmt.Scan(&k)
	a := k / 3600
	b := k % 3600 / 60
	fmt.Printf("It is %d hours %d minutes.", a, b)
}
