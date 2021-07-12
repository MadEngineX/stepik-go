package main
import "fmt"

func main() {
	var n, b int
	var a []int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		a = append(a, b)
	}
	fmt.Print(a[3])
}