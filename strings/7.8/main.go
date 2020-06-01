package main

import (
	"fmt"
	"strconv"
	"strings"
)

// look and say series
// 1, one 1, two 1, one 2 one 1, one 1 one 2 two 1 .....
// 1, 11, 21, 1211, 111221, 312211, 13112221, 1113213211
func main() {
	input := 10
	var result []string

	result = append(result, "1")

	for i := 1; i <= input; i++ {
		toProcess := strings.Split(result[i-1], "")
		count := 0
		temp := ""
		res := ""

		for indx, val := range toProcess {
			if indx == 0 {
				count++
				temp = val
			}

			if indx != 0 && temp == val {
				count++
			}

			if indx != 0 && temp != val {
				res = res + strconv.Itoa(count) + toProcess[indx-1]
				temp = val
				count = 1
			}
		}

		res = res + strconv.Itoa(count) + toProcess[len(toProcess)-1]
		result = append(result, res)
	}

	fmt.Println(result)
}
