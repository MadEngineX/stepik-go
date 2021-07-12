package main

import "fmt"

func main() {
	var input_var int
	fmt.Scan(&input_var)
	a := input_var%10
	b := input_var/10%10
	c := input_var/100%10
	fmt.Print(a+b+c)
}
