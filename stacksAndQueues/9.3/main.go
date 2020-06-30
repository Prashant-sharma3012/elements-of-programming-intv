package main

import (
	"fmt"
	"strings"
)

// test if parenthesis are good means every opening parenthesis has a closing
// allowed => ( ) { } [ ]
// best way of doing this is using stacks
// we will be using a slice to do this
func main() {
	var openingParens []string
	input := "({[()]})[]{}()"
	inputSplitted := strings.Split(input, "")

	for _, val := range inputSplitted {
		if val == "(" || val == "{" || val == "[" {
			openingParens = append(openingParens, val)
		} else {
			if (val == ")" && openingParens[len(openingParens)-1] == "(") ||
				(val == "}" && openingParens[len(openingParens)-1] == "{") ||
				(val == "]" && openingParens[len(openingParens)-1] == "[") {
				openingParens = openingParens[:len(openingParens)-1]
			} else {
				break
			}
		}
	}

	if len(openingParens) > 0 {
		fmt.Println("Not a valid expression")
	} else {
		fmt.Println("Valid expression")
	}
}
