package main

import (
	"fmt"
	"strings"
)

// convert roman numeral to number

var numberToRomaMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main() {
	input := "LIX"
	var res int
	splitted := strings.Split(input, "")

	for indx, val := range splitted {
		if indx == 0 {
			res = res + numberToRomaMap[val]
		} else {
			if numberToRomaMap[splitted[indx]] >= numberToRomaMap[splitted[indx-1]] {
				res = res + numberToRomaMap[val]
			} else {
				res = res - numberToRomaMap[val]
			}
		}
	}

	fmt.Println(res)
}
