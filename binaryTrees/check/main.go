package main

import (
	"fmt"

	"github.com/elements-of-programming-intv/binaryTrees/bstree"
)

func main() {

	elements := []int{30, 80, 60, 10, 20, 70, 25, 90, 100, 40, 5}

	t := bstree.GetBTreeInt(50)

	for _, val := range elements {
		t.Insert(val)
	}

	fmt.Println(t.InOrder())
	fmt.Println(t.PostOrder())
	fmt.Println(t.PreOrder())

	t.Remove(20)

	fmt.Println(t.InOrder())
	fmt.Println(t.PostOrder())
	fmt.Println(t.PreOrder())

	fmt.Println(t.Has(80))
}
