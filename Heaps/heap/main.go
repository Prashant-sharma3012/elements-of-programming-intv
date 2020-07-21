package heap

import "fmt"

var MinHeap []int
var MaxHeap []int

// parent = i
// childer = 2i & 2i + 1 = (7) => 3

// [0, 1,2, 3,4, 5,6, 7,8]

func AddToMinHeap(val int) []int {
	MinHeap = append(MinHeap, val)

	if len(MinHeap) == 1 {
		return MinHeap
	} else {
		currentPos := len(MinHeap) - 1
		parentPost := 0
		// buble up the newly added element
		for currentPos != 0 {
			if currentPos%2 == 0 {
				parentPost = (currentPos - 2) / 2
			} else {
				parentPost = (currentPos - 1) / 2
			}

			if MinHeap[parentPost] > MinHeap[currentPos] {
				temp := MinHeap[parentPost]
				MinHeap[parentPost] = MinHeap[currentPos]
				MinHeap[currentPos] = temp
				currentPos = parentPost
			} else {
				break
			}
		}
	}

	return MinHeap
}

func GetLowestFromMinHeap() int {
	min := MinHeap[0]
	MinHeap = append(MinHeap[len(MinHeap)-1:], MinHeap[1:len(MinHeap)-1]...)

	// since we are removing the minimun need to find new root
	currePos := 0

	for {
		// check if node has both children
		if currePos*2+1 >= len(MinHeap) {
			break
		}

		if currePos*2+2 >= len(MinHeap) {
			if MinHeap[currePos] > MinHeap[currePos*2+1] {
				temp := MinHeap[currePos*2+1]
				MinHeap[currePos*2+1] = MinHeap[currePos]
				MinHeap[currePos] = temp
				currePos = currePos*2 + 1
			} else {
				break
			}
		}

		if currePos*2+1 < len(MinHeap) && currePos*2+2 < len(MinHeap) {
			if MinHeap[currePos] > MinHeap[currePos*2+2] || MinHeap[currePos] > MinHeap[currePos*2+1] {
				if MinHeap[currePos*2+2] < MinHeap[currePos*2+1] {
					temp := MinHeap[currePos*2+2]
					MinHeap[currePos*2+2] = MinHeap[currePos]
					MinHeap[currePos] = temp
					currePos = currePos*2 + 2
				} else {
					temp := MinHeap[currePos*2+1]
					MinHeap[currePos*2+1] = MinHeap[currePos]
					MinHeap[currePos] = temp
					currePos = currePos*2 + 1
				}
			} else {
				break
			}
		}
	}

	return min
}

func AddToMaxHeap(val int) []int {
	MaxHeap = append(MaxHeap, val)

	if len(MaxHeap) == 1 {
		return MaxHeap
	} else {
		currentPos := len(MaxHeap) - 1
		parentPost := 0
		// buble up the newly added element
		for currentPos != 0 {
			if currentPos%2 == 0 {
				parentPost = (currentPos - 2) / 2
			} else {
				parentPost = (currentPos - 1) / 2
			}

			if MaxHeap[parentPost] < MaxHeap[currentPos] {
				temp := MaxHeap[parentPost]
				MaxHeap[parentPost] = MaxHeap[currentPos]
				MaxHeap[currentPos] = temp
				currentPos = parentPost
			} else {
				break
			}
		}
	}

	return MaxHeap
}

func GetHighestFromMaxHeap() int {
	max := MaxHeap[0]
	MaxHeap = append(MaxHeap[len(MaxHeap)-1:], MaxHeap[1:len(MaxHeap)-1]...)

	// since we are removing the max need to find new root

	currePos := 0

	for {
		// check if node has both children
		if currePos*2+1 >= len(MaxHeap) {
			break
		}

		if currePos*2+2 >= len(MaxHeap) {
			if MaxHeap[currePos] < MaxHeap[currePos*2+1] {
				temp := MaxHeap[currePos*2+1]
				MaxHeap[currePos*2+1] = MaxHeap[currePos]
				MaxHeap[currePos] = temp
				currePos = currePos*2 + 1
			} else {
				break
			}
		}

		if currePos*2+1 < len(MaxHeap) && currePos*2+2 < len(MaxHeap) {
			if MaxHeap[currePos] < MaxHeap[currePos*2+2] || MaxHeap[currePos] < MaxHeap[currePos*2+1] {
				if MaxHeap[currePos*2+2] > MaxHeap[currePos*2+1] {
					temp := MaxHeap[currePos*2+2]
					MaxHeap[currePos*2+2] = MaxHeap[currePos]
					MaxHeap[currePos] = temp
					currePos = currePos*2 + 2
				} else {
					temp := MaxHeap[currePos*2+1]
					MaxHeap[currePos*2+1] = MaxHeap[currePos]
					MaxHeap[currePos] = temp
					currePos = currePos*2 + 1
				}
			} else {
				break
			}
		}
	}
	return max
}

func PrintMin() {
	fmt.Println(MinHeap)
}

func PrintMax() {
	fmt.Println(MaxHeap)
}
