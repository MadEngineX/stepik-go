package main

import "fmt"

func main() {
	var a, b, max int
	fmt.Scan(&a, &b)

	for i := b; i >= a; i-- {
		if i % 7 == 0 {
			max = i
			break
		}
	}
	if max != 0 {
		fmt.Print(max)
	} else if b == 0 {
		fmt.Print(0)
	} else {
		fmt.Print("NO")
	}
}

