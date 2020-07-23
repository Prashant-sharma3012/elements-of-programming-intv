package main

import (
	"fmt"

	"github.com/elements-of-programming-intv/Heaps/heap"
)

// find median of a series of numbers

func main() {
	// for sake of brevity ew wil lsotrethe input in an array
	input := []int{1, 0, 3, 5, 2, 0, 1}

	// Min heap should store larger values
	// max heap shouldsotre smaller values
	// they should be balanced
	for indx, val := range input {
		if indx == 0 {
			heap.AddToMinHeap(val)
		} else {
			if len(heap.MinHeap) > len(heap.MaxHeap) {
				if heap.MinHeap[0] < val {
					temp := heap.GetLowestFromMinHeap()
					heap.AddToMinHeap(val)
					heap.AddToMaxHeap(temp)
				} else {
					heap.AddToMaxHeap(val)
				}
			} else {
				if heap.MaxHeap[0] > val {
					temp := heap.GetHighestFromMaxHeap()
					heap.AddToMaxHeap(val)
					heap.AddToMinHeap(temp)
				} else {
					heap.AddToMinHeap(val)
				}
			}
		}

		if indx == 0 {
			fmt.Print("Median ==> ")
			fmt.Println(heap.MinHeap[0])
		} else {
			if len(heap.MinHeap) == len(heap.MaxHeap) {
				fmt.Print("Median ==> ")
				fmt.Println(float64(heap.MinHeap[0]+heap.MaxHeap[0]) / float64(2))
			} else {
				if len(heap.MinHeap) > len(heap.MaxHeap) {
					fmt.Print("Median ==> ")
					fmt.Println(heap.MinHeap[0])
				} else {
					fmt.Print("Median ==> ")
					fmt.Println(heap.MaxHeap[0])
				}
			}
		}
	}
}
