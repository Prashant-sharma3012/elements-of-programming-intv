package main

import (
	"fmt"

	"github.com/elements-of-programming-intv/binaryTrees/bstree"
)

var parent map[int]*bstree.BSTreeInt

func nonRecursiceInOrder(t *bstree.BSTreeInt) []int {
	var inorder []int
	var curr, prev, next *bstree.BSTreeInt

	if t.Left != nil {
		curr = t.Left
		prev = t
		parent[t.Left.Val] = t
	} else {
		if t.Right != nil {
			curr = t.Right
			prev = t
			parent[t.Right.Val] = t
		}
	}

	if _, ok := parent[curr.Val]; ok {
		// going down
		if parent[curr.Val] == prev {
			if curr.Left != nil {
				next = curr.Left
				parent[curr.Left.Val] = prev
			} else {
				inorder = append(inorder, curr.Val)

				if curr.Right != nil {
					next = curr.Right
				} else {
					next = parent[curr.Val]
				}
			}
		} else {
			if curr.Left.Val == prev.Val {
				inorder = append(inorder, curr.Val)

				if curr.Right != nil {
					next = curr.Right
				} else {
					next = parent[curr.Val]
				}
			} else {
				next = parent[curr.Val]
			}
		}

		prev = curr
		curr = next

	}

	return inorder
}

func main() {
	elements := []int{30, 80, 60, 10, 20, 70, 25, 90, 100, 40, 5}

	t := bstree.GetBTreeInt(50)

	for _, val := range elements {
		t.Insert(val)
	}

	// inorder traversal with recirsion
	fmt.Println(t.InOrder())

	// inorder with o(1) space
	// for this all we need is store parent node on each node
	// instead i will do it without recirsion and o(h) complexity (h is height of tree)
	// will be using a Map to store parent node of a node
	fmt.Println(nonRecursiceInOrder(t))
}
