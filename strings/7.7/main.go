package main

import (
	"fmt"
	"strconv"
	"strings"
)

// take input a phone number (0 and 1 should not be in the number) and print all albhabet variants
// old phones had alphabets ofr each number 3 or 4

var result []string
var phoneEncoding = []string{"0", "1", "ABC", "DEF", "GHI", "JKL", "MNO", "PQRS", "TUV", "WXYZ"}

func getCombinations(input []string) {
	var temp []string
	// if all strings are of length 1
	// [A, G, M] merge and push to result
	for _, val := range input {
		if len(val) == 1 {
			temp = append(temp, val)
		}
	}

	if len(temp) == len(input) {
		result = append(result, strings.Join(temp, ""))
		return
	}

	for indx, val := range input {
		if len(val) > 1 {
			var path1, path2, temp1, temp2 []string

			// split and call again
			parts := strings.SplitN(val, "", 2)
			temp1 = append([]string{parts[0]}, input[indx+1:]...)
			path1 = append(input[:indx], temp1...)

			getCombinations(path1)

			temp2 = append([]string{parts[1]}, input[indx+1:]...)
			path2 = append(input[:indx], temp2...)
			getCombinations(path2)
		}
	}
}

func main() {

	phoneNo := "23"
	phonenoToArray := strings.Split(phoneNo, "")
	var encodedVal []string
	var at int

	for _, val := range phonenoToArray {
		at, _ = strconv.Atoi(val)
		encodedVal = append(encodedVal, phoneEncoding[at])
	}

	getCombinations(encodedVal)

	fmt.Println(result)
}