package main

import (
	"fmt"

	"github.com/elements-of-programming-intv/binaryTrees/bstree"
)

func loadRightSibling(b *bstree.BSTreeInt) {
	start := b

	if start != nil && start.Left != nil {
		loadLowerLevel(start)
		start = start.Left
	}
}

func loadLowerLevel(b *bstree.BSTreeInt) {
	start := b
	if start != nil {
		start.Left.RightSibling = start.Right

		if start.RightSibling != nil {
			start.Right.RightSibling = start.RightSibling.Left
		}

		start = start.RightSibling
	}
}

func main() {
	elements := []int{250, 100, 300, 50, 150, 270, 350, 800, 700, 900, 600, 750, 850, 950}

	t := bstree.GetBTreeInt(500)

	for _, val := range elements {
		t.Insert(val)
	}

	fmt.Println(t.PreOrder())

	loadRightSibling(t)
}
