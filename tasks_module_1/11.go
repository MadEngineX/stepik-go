package main

import "fmt"

func main() {
	var n uint8
	fmt.Scan(&n)
	if (n == 1) || (n % 10 == 1) && (n != 11) {
		fmt.Printf("%d korova", n)
	} else if ((n > 1 && n < 5) || (n % 10 > 1 && n % 10 < 5)) && (n != 12) && (n != 13) && (n != 14) {
		fmt.Printf("%d korovy", n)
	} else {
		fmt.Printf("%d korov", n)
	}

}
