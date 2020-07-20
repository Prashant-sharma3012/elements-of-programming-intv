package bstree

import (
	"errors"
	"fmt"
)

type BSTreeInt struct {
	Val          int
	Left         *BSTreeInt
	Right        *BSTreeInt
	RightSibling *BSTreeInt
}

func GetBTreeInt(val int) *BSTreeInt {
	return &BSTreeInt{
		Val:          val,
		Left:         nil,
		Right:        nil,
		RightSibling: nil,
	}
}

func addToTree(val int, b *BSTreeInt) {
	if b.Val > val {
		if b.Left == nil {
			b.Left = &BSTreeInt{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			return
		}
		addToTree(val, b.Left)
		return
	} else {
		if b.Right == nil {
			b.Right = &BSTreeInt{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			return
		}

		addToTree(val, b.Right)
		return
	}
}

func find(val int, b *BSTreeInt) bool {
	if b.Val == val {
		return true
	} else {
		if b.Val > val {
			return find(val, b.Left)
		}
		return find(val, b.Right)
	}

	return false
}

func getNode(val int, node *BSTreeInt, parent *BSTreeInt) (*BSTreeInt, *BSTreeInt) {
	if node.Val == val {
		return node, parent
	}

	if node.Val > val {
		return getNode(val, node.Left, node)
	}

	return getNode(val, node.Right, node)
}

func getInOrderSuccesor(val int, root *BSTreeInt) (*BSTreeInt, error) {
	inOrderArr := root.InOrder()
	for indx, k := range inOrderArr {
		if k == val {
			node, _ := getNode(inOrderArr[indx+1], root, root)
			return node, nil
		}
	}

	return nil, errors.New("Not Found")
}

func removeFromTree(val int, b *BSTreeInt) {
	node, parent := getNode(val, b, b)

	// check if its a leaf node
	if node.Left == nil && node.Right == nil {
		fmt.Println("Is leaf Node")
		if parent.Left.Val == node.Val {
			parent.Left = nil
			return
		}
		parent.Right = nil
		return
	}

	// if it has one child
	if (node.Left == nil && node.Right != nil) ||
		(node.Left != nil && node.Right == nil) {
		var temp *BSTreeInt
		if node.Left != nil {
			temp = node.Left
		} else {
			temp = node.Right
		}

		fmt.Println(temp)

		if parent.Left.Val == node.Val {
			parent.Left = temp
			return
		}
		parent.Right = temp
		return

	}

	// if node has two children
	successor, err := getInOrderSuccesor(val, b)
	if err != nil {
		b = node.Left
		return
	}

	// if node is not root
	if parent.Val != node.Val {
		if parent.Left.Val == node.Val {
			parent.Left = successor
			return
		}
		parent.Right = successor
	}

	successor.Left = node.Left
	successor.Right = node.Right
}

func (b *BSTreeInt) Insert(val int) {
	addToTree(val, b)
}

func (b *BSTreeInt) Has(val int) bool {
	return find(val, b)
}

func (b *BSTreeInt) GetNodeByVal(val int) *BSTreeInt {
	node, _ := getNode(val, b, b)
	return node
}

func (b *BSTreeInt) Remove(val int) {
	removeFromTree(val, b)
}

var inOrder []int

func traverseInOrder(b *BSTreeInt) []int {
	if b.Left != nil {
		traverseInOrder(b.Left)
	}

	inOrder = append(inOrder, b.Val)

	if b.Right != nil {
		traverseInOrder(b.Right)
	}

	return inOrder
}

// Left Root Right
func (b *BSTreeInt) InOrder() []int {
	inOrder = []int{}
	return traverseInOrder(b)
}

var preOrder []int

func traversePreOrder(b *BSTreeInt) []int {

	preOrder = append(preOrder, b.Val)

	if b.Left != nil {
		traversePreOrder(b.Left)
	}

	if b.Right != nil {
		traversePreOrder(b.Right)
	}

	return preOrder
}

// Root Left Right
func (b *BSTreeInt) PreOrder() []int {
	return traversePreOrder(b)
}

var postOrder []int

func traversePostOrder(b *BSTreeInt) []int {
	if b.Left != nil {
		traversePostOrder(b.Left)
	}

	if b.Right != nil {
		traversePostOrder(b.Right)
	}

	postOrder = append(postOrder, b.Val)

	return postOrder
}

// Left Right Root
func (b *BSTreeInt) PostOrder() []int {
	return traversePostOrder(b)
}

func (b *BSTreeInt) LevelOrder() {

}

func CreateTreeFromInorder(arr []int) {

}
