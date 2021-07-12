package main

import "fmt"

func main () {
	var input_var int
	fmt.Scan(&input_var)
	a := input_var%10
	b := input_var/10%10
	c := input_var/100%10
	d := input_var/1000%10
	e := input_var/10000%10
	f := input_var/100000


	if  (a + b + c) == (d + e + f) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
