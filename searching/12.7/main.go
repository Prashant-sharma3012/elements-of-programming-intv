package main

import "fmt"

var data = []int{3, 2, 5, 1, 1, 5, 2, 4}

func main() {

	// brute force approach
	// 2(n - 1) comparison
	// min := data[0]
	// max := data[0]

	// for _, val := range data {
	// 	if val < min {
	// 		min = val
	// 	} else {
	// 		if val > max {
	// 			max = val
	// 		}
	// 	}
	// }

	// A better approach with 3n/2 -1 comparisons
	min := data[0]
	max := data[1]

	if data[0] > data[1] {
		min = data[1]
		max = data[0]
	}

	for i := 2; i < len(data); i = i + 2 {
		tempMin := 0
		tempMax := 0
		if data[i] > data[i+1] {
			tempMin = data[i+1]
			tempMax = data[i]
		} else {
			tempMin = data[i]
			tempMax = data[i+1]
		}

		if min > tempMin {
			min = tempMin
		}

		if max < tempMax {
			max = tempMax
		}
	}

	// if lenght is even we have to compare the last element as well
	if len(data)%2 == 0 {
		last := data[len(data)-1]
		if min > last {
			min = last
		} else {
			if max < last {
				max = last
			}
		}
	}

	fmt.Println("min ===> ", min)
	fmt.Println("max ===> ", max)
}
