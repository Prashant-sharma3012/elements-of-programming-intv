package main

import (
	"fmt"

	"github.com/elements-of-programming-intv/Heaps/heap"
)

func main() {
	heapdata := []int{30, 20, 10, 5, 15, 45, 11, 25, 35}

	for _, val := range heapdata {
		heap.AddToMinHeap(val)
		heap.AddToMaxHeap(val)
	}

	fmt.Println(heapdata)
	heap.PrintMin()
	heap.PrintMax()

	fmt.Println(heap.GetHighestFromMaxHeap())
	fmt.Println(heap.GetLowestFromMinHeap())

	heap.PrintMin()
	heap.PrintMax()
}
