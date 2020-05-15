package main

import "fmt"

// remove duplicates from a sorted array in o(n) time and 0(1) space
// [2,3,4,4,5,5,7,7,9,11] => [2,3,4,5,7,9,11,0,0,0]
func main() {
	input := []int{2, 3, 4, 4, 4, 4, 5, 5, 5, 5, 7, 7, 7, 7, 9, 11}

	currentNum := input[0]
	currentPos := 0

	for indx := range input {
		if indx == len(input)-1 {
			break
		}

		if currentNum == input[indx+1] {
			input[indx+1] = 0
			if input[indx] != 0 {
				currentPos = indx + 1
			}
		} else {
			currentNum = input[indx+1]

			if currentPos != 0 {
				input[currentPos] = input[indx+1]
				input[indx+1] = 0
				currentPos++
			}
		}
	}

	fmt.Println(input)
}
