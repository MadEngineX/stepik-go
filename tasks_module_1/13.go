package main

import "fmt"

func main() {
	//1 1 2 3 5 8
	var a, count int
	var  fibo []int
	fmt.Scan(&a)

	fibo = append(fibo, 1, 1)

	for i:=2;  i < a; i++ {
		fibo = append(fibo, fibo[i-1]+fibo[i-2])
		if fibo[i] > a {
			break
		}
	}
	for _, elem := range fibo {
		count++
		if elem == a {
			fmt.Print(count)
			break
		}
	}
	//fmt.Print(fibo)
	if count == len(fibo) && fibo[len(fibo)-1] != a {
		fmt.Print(-1)
	}
}
