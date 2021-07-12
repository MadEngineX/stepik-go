package main

import "fmt"

func main() {
    var a float64
    fmt.Scan(&a)
    b := a*a
    if a <= 0 {
    	fmt.Printf("число %.2f не подходит", a)
	} else if b > 10000 {
		fmt.Printf("%e", a)
	} else {
		fmt.Printf("%.4f", b)
	}

}
