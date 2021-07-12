package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	strings.TrimSuffix(text, "\n")
	textRunes := []rune(text)
	if unicode.IsUpper(textRunes[0])  && textRunes[len(textRunes) - 2] == '.' {
		fmt.Print("Right")
	} else {
	fmt.Print("Wrong")
	}
}
