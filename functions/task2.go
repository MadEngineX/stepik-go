package main

import "fmt"

func main(){
	fmt.Print(fibonacci(4))
}
func fibonacci(n int) int{
	var  fibo []int

	fibo = append(fibo, 1, 1)

	for i:=2;  i < n; i++ {
		fibo = append(fibo, fibo[i-1]+fibo[i-2])
	}

	return fibo[n-1]
}