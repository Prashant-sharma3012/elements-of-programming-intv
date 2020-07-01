package bstree

import "errors"

type BSTreeInt struct {
	val   int
	left  *BSTreeInt
	right *BSTreeInt
}

func GetBTreeInt(val int) *BSTreeInt {
	return &BSTreeInt{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func addToTree(val int, b *BSTreeInt) {
	if b.val > val {
		if b.left == nil {
			b.left = &BSTreeInt{
				val:   val,
				left:  nil,
				right: nil,
			}
			return
		}
		addToTree(val, b.left)
		return
	} else {
		if b.right == nil {
			b.right = &BSTreeInt{
				val:   val,
				left:  nil,
				right: nil,
			}
			return
		}

		addToTree(val, b.right)
		return
	}
}

func find(val int, b *BSTreeInt) bool {
	if b.val == val {
		return true
	} else {
		if b.val > val {
			return find(val, b.left)
		}
		return find(val, b.right)
	}

	return false
}

func getNode(val int, node *BSTreeInt, parent *BSTreeInt) (*BSTreeInt, *BSTreeInt) {
	if node.val == val {
		return node, parent
	}

	if node.val > val {
		return getNode(val, node.left, node)
	}

	return getNode(val, node.right, node)
}

func getInOrderSuccesor(val int, root *BSTreeInt) (*BSTreeInt, error) {
	inOrderArr := root.InOrder()
	for _, k := range inOrderArr {
		if k == val {
			node, _ := getNode(val, root, root)
			return node, nil
		}
	}

	return nil, errors.New("Not Found")
}

func removeFromTree(val int, b *BSTreeInt) {
	node, _ := getNode(val, b, b)

	successor, err := getInOrderSuccesor(val, b)
	if err != nil {
		b = node.left
		return
	}

	successor.left = node.left
	successor.right = node.right
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
	if b.left != nil {
		traverseInOrder(b.left)
	}

	inOrder = append(inOrder, b.val)

	if b.right != nil {
		traverseInOrder(b.right)
	}

	return inOrder
}

// Left Root Right
func (b *BSTreeInt) InOrder() []int {
	return traverseInOrder(b)
}

var preOrder []int

func traversePreOrder(b *BSTreeInt) []int {

	preOrder = append(preOrder, b.val)

	if b.left != nil {
		traversePreOrder(b.left)
	}

	if b.right != nil {
		traversePreOrder(b.right)
	}

	return preOrder
}

// Root Left Right
func (b *BSTreeInt) PreOrder() []int {
	return traversePreOrder(b)
}

var postOrder []int

func traversePostOrder(b *BSTreeInt) []int {
	if b.left != nil {
		traversePostOrder(b.left)
	}

	if b.right != nil {
		traversePostOrder(b.right)
	}

	postOrder = append(postOrder, b.val)

	return postOrder
}

// Left Right Root
func (b *BSTreeInt) PostOrder() []int {
	return traversePostOrder(b)
}

func (b *BSTreeInt) LevelOrder() {

}
