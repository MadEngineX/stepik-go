package main

import "fmt"

func main() {
	var a uint64
	var b uint64
	fmt.Scan(&a)

    b = 1 + ((a - 1) % 9)
	fmt.Print(b)
}
