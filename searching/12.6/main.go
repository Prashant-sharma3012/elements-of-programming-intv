package main

import "fmt"

var data = [][]int{
	{-1, 2, 4, 4, 6},
	{1, 5, 5, 9, 21},
	{3, 6, 6, 9, 22},
	{3, 6, 8, 10, 24},
	{6, 8, 9, 12, 25},
	{8, 10, 12, 13, 40},
}

// find a number in a 2d array
func main() {
	search := 12
	start := 0
	end := len(data[0]) - 1
	found := false

	for start < len(data) && end > -1 {
		if data[start][end] == search {
			found = true
			fmt.Println("found")
			break
		}

		if data[start][end] < search {
			start++
		} else {
			end--
		}
	}

	if !found {
		fmt.Println("not found")
	}
}
