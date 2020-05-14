package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Multiply to long numbers that come in input as arrays
// for brevity taking small numbers
// example - [4,7,6], [9,4,3] = [4,4,8,8,6,8]
func main() {
	var temp, lastPos int

	num1 := []int{9, 9, 9}
	num2 := []int{9, 9, 9}
	num3 := make([]int, len(num1)+len(num2))

	for i, val1 := range num1 {
		for j, val2 := range num2 {

			res := val1*val2 + temp + num3[i+j]

			if res > 9 {
				conv := strings.Split(strconv.Itoa(res), "")
				res, _ = strconv.Atoi(conv[len(conv)-1])
				temp, _ = strconv.Atoi(strings.Join(conv[:len(conv)-1], ""))
			} else {
				temp = 0
			}

			num3[i+j] = res
			lastPos = i + j
		}
		if temp > 0 {
			num3[lastPos+1] = temp
			temp = 0
		}
	}

	fmt.Println(num3)
}
