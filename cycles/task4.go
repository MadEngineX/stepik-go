package main

import "fmt"

func main () {

	var n, max, count int
	// считываем числа пока не будет введен 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if n > max {
			max = n
		} else if n == max {
			count++
		}
	}

	fmt.Println(count)
}