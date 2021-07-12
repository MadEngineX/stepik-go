package main

import "fmt"

func main() {
	test("I like Go!")
}
func test(text string) {
	fmt.Println(string(text[2]))
}


