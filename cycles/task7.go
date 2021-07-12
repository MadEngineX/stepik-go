package main

import "fmt"

func main() {
	var x, p, y, i, count int
	fmt.Scan(&x, &p, &y)
	for i = x * 100; i < y * 100; {
		i = i + i/100*p
		count++
	}
	fmt.Println(count)
}
