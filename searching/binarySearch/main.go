package main

import "fmt"

var data = []int{1, 4, 7, 11, 12, 13, 40, 60}

func main() {

	search := 2
	start := 0
	end := len(data) - 1
	found := false
	center := 0

	for start != end || found {
		if len(data)%2 == 0 {
			center = (start + end + 1) / 2
		} else {
			center = (start + end) / 2
		}

		fmt.Println(center)

		if search == data[center] {
			found = true
			fmt.Println("Found")
			break
		}

		if search > data[center] {
			start = center + 1
		} else {
			end = center - 1
		}
	}
}
