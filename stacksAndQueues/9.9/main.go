package main

import "github.com/elements-of-programming-intv/stacksAndQueues/stack"

var mainStack = stack.MakeStack()
var mocksQStack = stack.MakeStack()

// implementing queue using stack
type queue struct{}

// add to queue
func (q *queue) enQ(val interface{}) {
	mainStack.Push(val)
}

// remove from queue
func (q *queue) dQ() interface{} {
	if !mocksQStack.IsEmpty() {
		return mocksQStack.Pop()
	}

	for !mainStack.IsEmpty() {
		mocksQStack.Push(mainStack.Pop())
	}

	return mocksQStack.Pop()
}

func main() {}
