package main

import (
	"fmt"
	"strings"
)

// snake string
// hello_world =>  e_lhlowrdlo
//    e         _             l
// h    l    o     w      r       d
//        l           o

func main() {
	input := "Hello_World"
	splitted := strings.Split(input, "")
	result := ""
	printLine1 := ""
	printLine2 := ""
	printLine3 := ""

	diff := 0
	// top
	for i := 1; i < len(splitted); i = i + 4 {
		padding := strings.Repeat(" ", i-diff)
		diff = i + 1
		printLine1 = printLine1 + padding + splitted[i]
		result = result + splitted[i]
	}

	diff = 0

	// middle
	for i := 0; i < len(splitted); i = i + 2 {
		padding := strings.Repeat(" ", i-diff)
		diff = i + 1
		printLine2 = printLine2 + padding + splitted[i]
		result = result + splitted[i]
	}

	diff = 0

	// bottom
	for i := 3; i < len(splitted); i = i + 4 {
		padding := strings.Repeat(" ", i-diff)
		diff = i + 1
		printLine3 = printLine3 + padding + splitted[i]
		result = result + splitted[i]
	}

	fmt.Println(result)

	fmt.Println(printLine1)
	fmt.Println(printLine2)
	fmt.Println(printLine3)
}
